/*
Implementation of a Queue data structure (FIFO)
Based on the custom doubly linked list package of this repo (dll)
*/

package queuell

import (
	"fmt"

	"github.com/joselws/go-utils/linked_lists/dll"
)

// Custom queue data structure based on the doubly linked list
// located in the dll package of this repo.
type Queue[T comparable] struct {
	list *dll.LinkedList[T]
	size int
}

// If you pass size of zero (0), your queue can grow infinitely.
func NewQueue[T comparable](size int) *Queue[T] {
	return &Queue[T]{
		list: dll.NewDoublyLinkedList[T](),
		size: size,
	}
}

// Returns "Queue(length)" if the queue has no size (0),
// or "Queue(length/size)" if the queue has a valid size.
func (queue Queue[T]) String() string {
	if queue.size == 0 {
		return fmt.Sprintf("Queue(%v)", queue.list.Length)
	}
	return fmt.Sprintf("Queue(%v/%v)", queue.list.Length, queue.size)
}

// Returns the number of elements in the queue.
func (queue *Queue[T]) Len() int {
	return queue.list.Length
}

// Returns true if the queue reached its "size" capacity.
// Always returns false if the queue has no size (0).
func (queue *Queue[T]) IsFull() bool {
	if queue.size == 0 {
		return false
	}
	if queue.size == queue.list.Length {
		return true
	}
	return false
}

// Returns true if the number of elements in the queue is zero (0).
func (queue *Queue[T]) IsEmpty() bool {
	if queue.list.Length == 0 {
		return true
	}
	return false
}

// Returns true if the item was put in the queue.
// Returns false if the queue is already full
// and the item couldn't be enqueued into.
func (queue *Queue[T]) Enqueue(value T) bool {
	if queue.IsFull() {
		return false
	}
	queue.list.AppendRight(value)
	return true
}

// Take out the next item in the queue and return it.
// It returns false on ok if the queue is empty.
func (queue *Queue[T]) Dequeue() (T, bool) {
	var value T
	if queue.Len() == 0 {
		return value, false
	}
	value, err := queue.list.PopLeft()
	if err != nil {
		return value, false
	}
	return value, true
}

// Get the next item from the queue without removing it.
// It returns false on ok if the queue is empty.
func (queue *Queue[T]) PeekNext() (T, bool) {
	var value T
	if queue.IsEmpty() {
		return value, false
	}
	value = queue.list.Head.Value
	return value, true
}

// Get the last item from the queue without removing it.
// It returns false on ok if the queue is empty.
func (queue *Queue[T]) PeekLast() (T, bool) {
	var value T
	if queue.IsEmpty() {
		return value, false
	}
	value = queue.list.Tail.Value
	return value, true
}
