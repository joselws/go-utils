/*
Implementation of a Stack data structure (LIFO)
*/

package stack

import (
	"errors"
	"fmt"
	"reflect"
)

type Stack[T any] struct {
	Items []T
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
