/*
Implementation of a Stack data structure (LIFO)
*/

package stack

import (
	"container/list"
	"errors"
	"fmt"
)

type Stack[T any] struct {
	Elements *list.List
}

// Stack constructor
func NewStack[T any]() Stack[T] {
	return Stack[T]{Elements: list.New()}
}

// Get the current length of the stack
func (thisStack *Stack[T]) Len() int {
	return thisStack.Elements.Len()
}

// Add a new item into the stack
func (thisStack *Stack[T]) Push(item T) {
	thisStack.Elements.PushBack(item)
}

/*
Remove next item of the stack and return its value
Returns error if the stack is empty
*/
func (thisStack *Stack[T]) Pop() (T, error) {
	var nextElement T
	element := thisStack.Elements.Back()
	if element == nil {
		return nextElement, errors.New("Cannot Pop() from empty stack.")
	}
	nextElement, ok := element.Value.(T)
	if !ok {
		return nextElement, errors.New(fmt.Sprintf("Mismatch type of Stack type and %T", element.Value))
	}
	thisStack.Elements.Remove(element)
	return nextElement, nil
}
