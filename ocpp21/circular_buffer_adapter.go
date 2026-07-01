package ocpp21

import "github.com/johnmaddison/ocpp-go/internal/cbuf"

type CircularBuffer struct {
	buf *cbuf.CircularBuffer[Request]
}

func NewCircularBuffer(capacity int) *CircularBuffer {
	return &CircularBuffer{buf: cbuf.NewCircularBuffer[Request](capacity)}
}

func (cb *CircularBuffer) Add(item Request) {
	cb.buf.Add(item)
}

func (cb *CircularBuffer) FindByMessageID(messageID string) (Request, bool) {
	return cb.buf.Find(func(r Request) bool { return r.Call.MessageID == messageID })
}

func (cb *CircularBuffer) RemoveByMessageID(messageID string) (Request, bool) {
	return cb.buf.RemoveWhere(func(r Request) bool { return r.Call.MessageID == messageID })
}

func (cb *CircularBuffer) Size() int { return cb.buf.Size() }

func (cb *CircularBuffer) IsFull() bool { return cb.buf.IsFull() }

func (cb *CircularBuffer) IsEmpty() bool { return cb.buf.IsEmpty() }

func (cb *CircularBuffer) Clear() { cb.buf.Clear() }

func (cb *CircularBuffer) GetAll() []Request { return cb.buf.GetAll() }
