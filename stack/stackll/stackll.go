package stackll

import (
	"errors"
	"fmt"

	"github.com/joselws/go-utils/linked_lists/sll"
)

// Custom stack data structure based on the singly linked list
// located in the sll package of this repo.
type Stack[T comparable] struct {
	linkedList *sll.LinkedList[T]
	size       int
}

// If you pass size of zero (0), your stack can grow infinitely.
func NewStack[T comparable](size int) *Stack[T] {
	stack := Stack[T]{
		linkedList: sll.NewSingleLinkedList[T](),
		size:       size,
	}
	return &stack
}

// Returns "Stack(length)" if the stack has no size (0),
// or "Stack(length/size)" if the stack has a valid size.
func (stack Stack[T]) String() string {
	if stack.size == 0 {
		return fmt.Sprintf("Stack(%v)", stack.linkedList.Length)
	}
	return fmt.Sprintf("Stack(%v/%v)", stack.linkedList.Length, stack.size)
}

// Returns the number of elements in the stack.
func (stack *Stack[T]) Len() int {
	return stack.linkedList.Length
}

// Returns true if the stack reached its "size" capacity.
// Always returns false if the stack has no size (0).
func (stack *Stack[T]) IsFull() bool {
	if stack.size == 0 {
		return false
	}
	if stack.size == stack.linkedList.Length {
		return true
	}
	return false
}

// Returns true if the number of elements in the stack is zero (0).
func (stack *Stack[T]) IsEmpty() bool {
	if stack.linkedList.Length == 0 {
		return true
	}
	return false
}

// Returns true if the item was put into the stack.
// Returns false if the stack is already full
// and the item couldn't be pushed into.
func (stack *Stack[T]) Push(value T) bool {
	if stack.IsFull() {
		return false
	}
	stack.linkedList.Prepend(value)
	return true
}

// Take out the next item in the stack and return it.
// It returns an error if the stack is empty.
func (stack *Stack[T]) Pop() (T, error) {
	var element T
	if stack.IsEmpty() {
		return element, errors.New("Cannot Pop from empty stack!")
	}
	element, err := stack.linkedList.PopFirst()
	if err != nil {
		return element, err
	}
	return element, nil
}

// Get the next item from the stack without removing it.
// It returns an error if the stack is empty.
func (stack *Stack[T]) PeekNext() (T, error) {
	var element T
	if stack.IsEmpty() {
		return element, errors.New("Cannot Peek() from empty stack!")
	}
	element = stack.linkedList.Head.Value
	return element, nil
}
