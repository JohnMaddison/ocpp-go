package cbuf

import "sync"

// CircularBuffer is a thread-safe circular buffer for generic items.
type CircularBuffer[T any] struct {
	buffer   []T
	head     int
	tail     int
	size     int
	capacity int
	full     bool
	mutex    sync.RWMutex
}

// NewCircularBuffer creates a new circular buffer with the specified capacity.
func NewCircularBuffer[T any](capacity int) *CircularBuffer[T] {
	return &CircularBuffer[T]{
		buffer:   make([]T, capacity),
		capacity: capacity,
		head:     0,
		tail:     0,
		size:     0,
		full:     false,
	}
}

// Add adds an item to the buffer. If the buffer is full, it overwrites the oldest item.
func (cb *CircularBuffer[T]) Add(item T) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.buffer[cb.tail] = item
	cb.tail = (cb.tail + 1) % cb.capacity

	if cb.full {
		// Buffer is full, move head forward (overwrite oldest)
		cb.head = (cb.head + 1) % cb.capacity
	} else {
		cb.size++
		if cb.tail == cb.head {
			cb.full = true
		}
	}
}

// Size returns the current number of items in the buffer.
func (cb *CircularBuffer[T]) Size() int {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.size
}

// IsFull returns true if the buffer is at capacity.
func (cb *CircularBuffer[T]) IsFull() bool {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.full
}

// IsEmpty returns true if the buffer is empty.
func (cb *CircularBuffer[T]) IsEmpty() bool {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	return cb.size == 0
}

// Clear removes all items from the buffer.
func (cb *CircularBuffer[T]) Clear() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.head = 0
	cb.tail = 0
	cb.size = 0
	cb.full = false
	var zero T
	for i := range cb.buffer {
		cb.buffer[i] = zero
	}
}

// GetAll returns a copy of all items in the buffer (from oldest to newest).
func (cb *CircularBuffer[T]) GetAll() []T {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()

	if cb.size == 0 {
		return []T{}
	}

	result := make([]T, cb.size)
	current := cb.head
	for i := 0; i < cb.size; i++ {
		result[i] = cb.buffer[current]
		current = (current + 1) % cb.capacity
	}
	return result
}

// FindIndex finds the index (0..size-1) of the first element matching pred.
// Returns -1 when not found.
func (cb *CircularBuffer[T]) FindIndex(pred func(T) bool) int {
	cb.mutex.RLock()
	defer cb.mutex.RUnlock()
	if cb.size == 0 {
		return -1
	}
	current := cb.head
	for i := 0; i < cb.size; i++ {
		if pred(cb.buffer[current]) {
			return i
		}
		current = (current + 1) % cb.capacity
	}
	return -1
}

// RemoveAt removes the item at logical index (0..size-1) and returns it.
// Returns ok=false when index is out of range.
func (cb *CircularBuffer[T]) RemoveAt(index int) (item T, ok bool) {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if index < 0 || index >= cb.size {
		return item, false
	}

	// Convert logical index to physical index
	phys := (cb.head + index) % cb.capacity
	item = cb.buffer[phys]

	// Shift items after phys backward
	cur := phys
	for j := 0; j < cb.size-1-index; j++ {
		next := (cur + 1) % cb.capacity
		cb.buffer[cur] = cb.buffer[next]
		cur = next
	}

	// Clear last slot and adjust tail/size
	var zero T
	cb.tail = (cb.tail - 1 + cb.capacity) % cb.capacity
	cb.buffer[cb.tail] = zero
	cb.size--
	cb.full = false
	return item, true
}
