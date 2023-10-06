/*
Implementation of a Stack data structure (LIFO)
Based on slices
*/

// TODO: compatible with Stack interface

package stackarr

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
)

type Stack[T any] struct {
	Items []T
	size  int
}

// Stack constructor
func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{size: size}
}

// Get the current length of the stack
func (stack *Stack[T]) Len() int {
	return len(stack.Items)
}

// Add a new item into the stack
func (stack *Stack[T]) Push(item T) bool {
	if stack.Len() == stack.size && stack.size != 0 {
		return false
	}
	stack.Items = append(stack.Items, item)
	return true
}

/*
Remove next item of the stack and return its value
Returns an error if the stack is empty
*/
func (stack *Stack[T]) Pop() (nextItem T, err error) {
	if stack.Len() == 0 {
		return nextItem, errors.New("Can't use Pop on an empty Stack")
	}
	var nextItemIndex int = stack.Len() - 1
	nextItem = stack.Items[nextItemIndex]
	stack.Items = stack.Items[:nextItemIndex]
	return nextItem, nil
}

/*
Extract many items from the stack in a new slice

args: amount of items to extract

returns: slice of the items extracted
for example, if the stack is [1, 2, 3], and you request 2, the result would be [2, 3]

the method returns an error if you request to extract more items than available
*/
func (stack *Stack[T]) ExtractMany(amountToExtract int) (items []T, err error) {
	var stackLength int = stack.Len()
	if stackLength < amountToExtract {
		return items, errors.New("Amount to extract exceeds stack size")
	}
	var startIndex int = stackLength - amountToExtract
	items = stack.Items[startIndex:stackLength]
	stack.Items = stack.Items[:startIndex]
	slices.Reverse(items)
	return items, nil
}

/*
Insert many items from a slice to the stack
For example, if the stack is [1] and the slice is [2, 3], the stack will be [1, 2, 3]
*/
func (stack *Stack[T]) InsertFromSlice(items []T) {
	for _, item := range items {
		stack.Push(item)
	}
}

/*
Returns the value of the next item without removing it from the stack
returns an error if the stack is empty
*/
func (stack *Stack[T]) PeekNext() (nextItem T, err error) {
	if stack.Len() == 0 {
		return nextItem, errors.New("Can't PeekNext on an empty Stack")
	}
	var nextItemIndex int = stack.Len() - 1
	nextItem = stack.Items[nextItemIndex]
	return nextItem, nil
}

// Prints Stack in format "Stack[T][item1 itemn]"
func (stack Stack[T]) String() string {
	stackType := reflect.TypeOf(stack.Items).Elem()
	return fmt.Sprintf("Stack[%v]%v", stackType, stack.Items)
}
