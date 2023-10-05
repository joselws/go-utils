/*
Implementation of a Stack data structure (LIFO)
*/

package stack

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
)

type Stack[T any] struct {
	Items []T
}

// Stack constructor
func NewStack[T any]() Stack[T] {
	return Stack[T]{}
}

// Get the current length of the stack
func (thisStack *Stack[T]) Length() int {
	return len(thisStack.Items)
}

// Add a new item into the stack
func (thisStack *Stack[T]) Push(item T) {
	thisStack.Items = append(thisStack.Items, item)
}

/*
Remove next item of the stack and return its value
Returns an error if the stack is empty
*/
func (thisStack *Stack[T]) Pop() (nextItem T, err error) {
	if thisStack.Length() == 0 {
		return nextItem, errors.New("Can't use Pop on an empty Stack")
	}
	var nextItemIndex int = thisStack.Length() - 1
	nextItem = thisStack.Items[nextItemIndex]
	thisStack.Items = thisStack.Items[:nextItemIndex]
	return nextItem, nil
}

/*
Extract many items from the stack in a new slice

args: amount of items to extract

returns: slice of the items extracted
for example, if the stack is [1, 2, 3], and you request 2, the result would be [2, 3]

the method returns an error if you request to extract more items than available
*/
func (thisStack *Stack[T]) ExtractMany(amountToExtract int) (items []T, err error) {
	var stackLength int = thisStack.Length()
	if stackLength < amountToExtract {
		return items, errors.New("Amount to extract exceeds stack size")
	}
	var startIndex int = stackLength - amountToExtract
	items = thisStack.Items[startIndex:stackLength]
	thisStack.Items = thisStack.Items[:startIndex]
	slices.Reverse(items)
	return items, nil
}

/*
Insert many items from a slice to the stack
For example, if the stack is [1] and the slice is [2, 3], the stack will be [1, 2, 3]
*/
func (thisStack *Stack[T]) InsertFromSlice(items []T) {
	for _, item := range items {
		thisStack.Push(item)
	}
}

/*
Returns the value of the next item without removing it from the stack
returns an error if the stack is empty
*/
func (thisStack *Stack[T]) InspectNextItem() (nextItem T, err error) {
	if thisStack.Length() == 0 {
		return nextItem, errors.New("Can't InspectNextItem on an empty Stack")
	}
	var nextItemIndex int = thisStack.Length() - 1
	nextItem = thisStack.Items[nextItemIndex]
	return nextItem, nil
}

// Prints Stack in format "Stack[T][item1 itemn]"
func (thisStack Stack[T]) String() string {
	stackType := reflect.TypeOf(thisStack.Items).Elem()
	return fmt.Sprintf("Stack[%v]%v", stackType, thisStack.Items)
}
