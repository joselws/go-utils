package deque

// Implementation of the deque interface
type Deque[T comparable] interface {
	IsFull() bool
	IsEmpty() bool
	Len() int
	AppendLeft(value T) bool
	AppendRight(value T) bool
	PeekLeft() (T, bool)
	PeekRight() (T, bool)
	PopLeft() (T, bool)
	PopRight() (T, bool)
}
