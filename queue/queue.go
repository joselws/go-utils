package queue

import (
	"errors"

	"github.com/joselws/go-utils/deque"
)

var ErrEmptyQueue = errors.New("queue is empty")
var ErrFullQueue = errors.New("queue is full")

type Queue[T any] struct {
	deque *deque.Deque[T]
}

// NewQueue creates a new Queue with the specified size.
// A size of 0 means the queue is unbounded.
func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{deque: deque.NewDeque[T](size)}
}

func (queue *Queue[T]) Len() int {
	return queue.deque.Len()
}

func (queue *Queue[T]) IsFull() bool {
	return queue.deque.IsFull()
}

func (queue *Queue[T]) IsEmpty() bool {
	return queue.deque.IsEmpty()
}

func (queue *Queue[T]) Enqueue(value T) error {
	err := queue.deque.AppendRight(value)
	if err != nil {
		return ErrFullQueue
	}
	return nil
}

func (queue *Queue[T]) Dequeue() (T, error) {
	value, err := queue.deque.PopLeft()
	if err != nil {
		return *new(T), ErrEmptyQueue
	}
	return value, nil
}

func (queue *Queue[T]) PeekNext() (T, error) {
	value, err := queue.deque.PeekLeft()
	if err != nil {
		return *new(T), ErrEmptyQueue
	}
	return value, nil
}

func (queue *Queue[T]) PeekLast() (T, error) {
	value, err := queue.deque.PeekRight()
	if err != nil {
		return *new(T), ErrEmptyQueue
	}
	return value, nil
}
