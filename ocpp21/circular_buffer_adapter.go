package ocpp21

import "github.com/johnmaddison/ocpp-go/internal/cbuf"

// CircularBuffer is a thin adapter over an internal generic buffer
// specialized for ocpp21 Request items.
type CircularBuffer struct {
	buf *cbuf.CircularBuffer[Request]
}

// NewCircularBuffer creates a new circular buffer with the specified capacity.
func NewCircularBuffer(capacity int) *CircularBuffer {
	return &CircularBuffer{buf: cbuf.NewCircularBuffer[Request](capacity)}
}

// Add adds an item to the buffer. If the buffer is full, it overwrites the oldest item.
func (cb *CircularBuffer) Add(item Request) {
	cb.buf.Add(item)
}

// FindByMessageID searches for an item by message ID.
func (cb *CircularBuffer) FindByMessageID(messageID string) (Request, bool) {
	return cb.buf.Find(func(r Request) bool { return r.Call.MessageID == messageID })
}

// RemoveByMessageID removes an item by message ID.
func (cb *CircularBuffer) RemoveByMessageID(messageID string) (Request, bool) {
	return cb.buf.RemoveWhere(func(r Request) bool { return r.Call.MessageID == messageID })
}

// Size returns the current number of items in the buffer.
func (cb *CircularBuffer) Size() int { return cb.buf.Size() }

// IsFull returns true if the buffer is at capacity.
func (cb *CircularBuffer) IsFull() bool { return cb.buf.IsFull() }

// IsEmpty returns true if the buffer is empty.
func (cb *CircularBuffer) IsEmpty() bool { return cb.buf.IsEmpty() }

// Clear removes all items from the buffer.
func (cb *CircularBuffer) Clear() { cb.buf.Clear() }

// GetAll returns a copy of all items in the buffer, from oldest to newest.
func (cb *CircularBuffer) GetAll() []Request { return cb.buf.GetAll() }
