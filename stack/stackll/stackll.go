package stackll

import (
	"errors"
	"fmt"
	"math"

	"github.com/joselws/go-utils/linked_lists/sll"
)

// TODO: docs
type Stack[T comparable] struct {
	linkedList *sll.LinkedList[T]
	size       float64
}

func NewStack[T comparable](size float64) *Stack[T] {
	var stackSize float64
	if size == 0 {
		stackSize = math.Inf(1)
	} else {
		stackSize = math.Floor(size)
	}
	stack := Stack[T]{
		linkedList: sll.NewSingleLinkedList[T](),
		size:       stackSize,
	}
	return &stack
}

func (stack Stack[T]) String() string {
	return fmt.Sprintf("Stack(%v/%v)", stack.linkedList.Length, stack.size)
}

func (stack *Stack[T]) IsFull() bool {
	if stack.size == float64(stack.linkedList.Length) {
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

func (stack *Stack[T]) Pop(value T) (T, error) {
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

func (stack *Stack[T]) PeekNext(value T) (T, error) {
	var element T
	if stack.IsEmpty() {
		return element, errors.New("Cannot Peek() from empty stack!")
	}
	element = stack.linkedList.Head.Value
	return element, nil
}
