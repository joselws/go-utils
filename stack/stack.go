/*
Implementation of a Stack data structure (LIFO)
Based on the standard library container/list
*/

package stack

import (
	"container/list"
	"errors"
	"fmt"
)

type Stack[T any] struct {
	elements *list.List
}

// Stack constructor.
func NewStack[T any]() Stack[T] {
	return Stack[T]{elements: list.New()}
}

// Get the current length of the stack.
func (thisStack *Stack[T]) Len() int {
	return thisStack.elements.Len()
}

// Add a new item into the stack.
func (thisStack *Stack[T]) Push(item T) {
	thisStack.elements.PushBack(item)
}

// Remove next item of the stack and return its value.
// Returns error if the stack is empty.
func (thisStack *Stack[T]) Pop() (T, error) {
	var nextElement T
	element := thisStack.elements.Back()
	if element == nil {
		return nextElement, errors.New("cannot Pop() from empty stack")
	}
	nextElement, ok := element.Value.(T)
	if !ok {
		return nextElement, fmt.Errorf("mismatch type of Stack type and %T", element.Value)
	}
	thisStack.elements.Remove(element)
	return nextElement, nil
}

func (thisStack *Stack[T]) PeekNext() (T, error) {
	var nextElement T
	element := thisStack.elements.Back()
	if element == nil {
		return nextElement, errors.New("cannot PeekNext() from empty stack")
	}
	nextElement, ok := element.Value.(T)
	if !ok {
		return nextElement, fmt.Errorf("mismatch type of Stack type and %T", element.Value)
	}
	return nextElement, nil
}
