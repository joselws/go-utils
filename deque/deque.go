/*
Implementation of a double-ended queue (deque).
Fast inserts and deletes at both ends of the structure.
Based on the doubly linked list structure
*/

package deque

import "github.com/joselws/go-utils/linked_lists/dll"

type Deque[T comparable] struct {
	list *dll.LinkedList[T]
	size int
}

func NewDeque[T comparable](size int) *Deque[T] {
	return &Deque[T]{
		list: dll.NewDoublyLinkedList[T](),
		size: size,
	}
}

func (deque *Deque[T]) Len() int {
	return deque.list.Length
}

func (deque *Deque[T]) IsFull() bool {
	if deque.size == 0 {
		return false
	}
	if deque.size == deque.list.Length {
		return true
	}
	return false
}

func (deque *Deque[T]) AppendRight(value T) bool {
	if deque.IsFull() {
		return false
	}
	deque.list.AppendRight(value)
	return true
}
