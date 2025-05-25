/*
Implementation of a Stack data structure (LIFO)
Based on the standard library container/list
*/

package stack

import (
	"errors"

	"github.com/joselws/go-utils/deque"
)

var ErrEmptyStack = errors.New("stack is empty")
var ErrFullStack = errors.New("stack is full")

type Stack[T any] struct {
	deque *deque.Deque[T]
}

// Stacks with size 0 are unbounded.
func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{deque: deque.NewDeque[T](size)}
}

// Get the current length of the stack.
func (stack *Stack[T]) Len() int {
	return stack.deque.Len()
}

// Add a new item into the stack.
func (stack *Stack[T]) Push(item T) error {
	err := stack.deque.AppendRight(item)
	if err != nil {
		return ErrFullStack
	}
	return nil
}

// Remove next item of the stack and return its value.
// Returns error if the stack is empty.
func (stack *Stack[T]) Pop() (T, error) {
	value, err := stack.deque.PopRight()
	if err != nil {
		return *new(T), ErrEmptyStack
	}
	return value, nil
}

func (thisStack *Stack[T]) PeekNext() (T, error) {
	value, err := thisStack.deque.PeekRight()
	if err != nil {
		return *new(T), ErrEmptyStack
	}
	return value, nil
}

func (stack *Stack[T]) IsEmpty() bool {
	return stack.deque.IsEmpty()
}

func (stack *Stack[T]) IsFull() bool {
	return stack.deque.IsFull()
}
