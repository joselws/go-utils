package stackll

import (
	"errors"
	"fmt"

	"github.com/joselws/go-utils/linked_lists/sll"
)

// TODO: docs
type Stack[T comparable] struct {
	linkedList *sll.LinkedList[T]
	size       int
}

func NewStack[T comparable](size int) *Stack[T] {
	stack := Stack[T]{
		linkedList: sll.NewSingleLinkedList[T](),
		size:       size,
	}
	return &stack
}

func (stack Stack[T]) String() string {
	if stack.size == 0 {
		return fmt.Sprintf("Stack(%v)", stack.linkedList.Length)
	}
	return fmt.Sprintf("Stack(%v/%v)", stack.linkedList.Length, stack.size)
}

func (stack *Stack[T]) Len() int {
	return stack.linkedList.Length
}

func (stack *Stack[T]) IsFull() bool {
	if stack.size == 0 {
		return false
	}
	if stack.size == stack.linkedList.Length {
		return true
	}
	return false
}

func (stack *Stack[T]) IsEmpty() bool {
	if stack.linkedList.Length == 0 {
		return true
	}
	return false
}

func (stack *Stack[T]) Push(value T) bool {
	if stack.IsFull() {
		return false
	}
	stack.linkedList.Prepend(value)
	return true
}

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

func (stack *Stack[T]) PeekNext() (T, error) {
	var element T
	if stack.IsEmpty() {
		return element, errors.New("Cannot Peek() from empty stack!")
	}
	element = stack.linkedList.Head.Value
	return element, nil
}
