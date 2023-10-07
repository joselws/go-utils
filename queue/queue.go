package queues

// Implementation of the Queue interface
type Queue[T comparable] interface {
	IsFull() bool
	IsEmpty() bool
	Len() int
	Enqueue(value T) bool
	Dequeue() (T, bool)
	PeekNext() (T, bool)
	PeekLast() (T, bool)
}
