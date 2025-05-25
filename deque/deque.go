/*
Implementation of a double-ended queue (deque).
Fast inserts and deletes at both ends of the structure.
Based on the doubly linked list structure
*/

package deque

import (
	"container/list"
	"errors"
)

var ErrEmptyDeque = errors.New("deque is empty")
var ErrFullDeque = errors.New("deque is full")

type Deque[T any] struct {
	list *list.List
	size int
}

// A 0 size deque is unbounded.
func NewDeque[T any](size int) *Deque[T] {
	return &Deque[T]{
		list: list.New(),
		size: size,
	}
}

func (deque *Deque[T]) Len() int {
	return deque.list.Len()
}

func (deque *Deque[T]) IsFull() bool {
	if deque.size == 0 {
		return false
	}
	if deque.size == deque.list.Len() {
		return true
	}
	return false
}

func (deque *Deque[T]) IsEmpty() bool {
	return deque.Len() == 0
}

func (deque *Deque[T]) AppendRight(value T) error {
	if deque.IsFull() {
		return ErrFullDeque
	}
	deque.list.PushBack(value)
	return nil
}
