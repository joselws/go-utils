package stacks

// Implementation of the Stack interface
type Stack[T comparable] interface {
	IsFull() bool
	IsEmpty() bool
	Len() int
	Push(value T) bool
	Pop() (T, error)
	PeekNext() (T, error)
}
