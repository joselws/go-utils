/*
Implementation of a double-ended queue (deque).
Fast inserts and deletes at both ends of the structure.
Based on the doubly linked list structure
*/

package deque

import (
	"container/list"
	"errors"
	"sync"
)

var ErrEmptyDeque = errors.New("deque is empty")
var ErrFullDeque = errors.New("deque is full")

type Deque[T any] struct {
	list *list.List
	size int
	mu   sync.Mutex
}

// A 0 size deque is unbounded.
func NewDeque[T any](size int) *Deque[T] {
	return &Deque[T]{
		list: list.New(),
		size: size,
	}
}

func (deque *Deque[T]) Len() int {
	deque.mu.Lock()
	defer deque.mu.Unlock()
	return deque.list.Len()
}

func (deque *Deque[T]) IsFull() bool {
	if deque.size == 0 {
		return false
	}
	return deque.size == deque.Len()
}

func (deque *Deque[T]) IsEmpty() bool {
	return deque.Len() == 0
}

func (deque *Deque[T]) AppendRight(value T) error {
	if deque.IsFull() {
		return ErrFullDeque
	}

	deque.mu.Lock()
	defer deque.mu.Unlock()

	deque.list.PushFront(value)
	return nil
}

func (deque *Deque[T]) PopRight() (T, error) {
	if deque.IsEmpty() {
		return *new(T), ErrEmptyDeque
	}

	deque.mu.Lock()
	defer deque.mu.Unlock()

	element := deque.list.Front()
	deque.list.Remove(element)
	return element.Value.(T), nil
}

func (deque *Deque[T]) AppendLeft(value T) error {
	if deque.IsFull() {
		return ErrFullDeque
	}

	deque.mu.Lock()
	defer deque.mu.Unlock()

	deque.list.PushBack(value)
	return nil
}

func (deque *Deque[T]) PopLeft() (T, error) {
	if deque.IsEmpty() {
		return *new(T), ErrEmptyDeque
	}

	deque.mu.Lock()
	defer deque.mu.Unlock()

	element := deque.list.Back()
	deque.list.Remove(element)
	return element.Value.(T), nil
}

func (deque *Deque[T]) PeekRight() (T, error) {
	if deque.IsEmpty() {
		return *new(T), ErrEmptyDeque
	}

	deque.mu.Lock()
	defer deque.mu.Unlock()

	element := deque.list.Front()
	return element.Value.(T), nil
}

func (deque *Deque[T]) PeekLeft() (T, error) {
	if deque.IsEmpty() {
		return *new(T), ErrEmptyDeque
	}

	deque.mu.Lock()
	defer deque.mu.Unlock()

	element := deque.list.Back()
	return element.Value.(T), nil
}
