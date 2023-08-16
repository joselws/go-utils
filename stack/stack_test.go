package stack

import (
	"fmt"
	"testing"
)

func TestStackCorrectlyInitialized(t *testing.T) {
	myStack := Stack[int]{}
	if len(myStack.Items) != 0 {
		t.Error("Stack length should be 0, not", len(myStack.Items))
	}
}

func TestPushMethod(t *testing.T) {
	myStack := Stack[int]{}
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	if len(myStack.Items) != 3 {
		t.Error("Stack length should be 3, not", len(myStack.Items))
	}
}

func TestPopMethod(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	lastItem, err := myStack.Pop()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if len(myStack.Items) != 2 {
		t.Error("Stack length should be 2, not", len(myStack.Items))
	}
	if lastItem != 300 {
		t.Error("Last Stack item should be 300, not", lastItem)
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	myStack := Stack[int]{}
	_, err := myStack.Pop()
	if err == nil {
		t.Error("Error should not be nil, but", err)
	}
}

func TestLength(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	var stackLength int = myStack.Length()
	if stackLength != 3 {
		t.Error("Stack length should be 3, not", stackLength)
	}
}

func TestString(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	var stackString string = fmt.Sprint(myStack)
	if stackString != "Stack[int][100 200 300]" {
		t.Error("Stack string representation should be Stack[int][100 200 300], not", stackString)
	}
}

func TestInspectNextItem(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	nextItem, err := myStack.InspectNextItem()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if len(myStack.Items) != 3 {
		t.Error("Stack length should be 3, not", len(myStack.Items))
	}
	if nextItem != 300 {
		t.Error("Last Stack item should be 300, not", nextItem)
	}
}

func TestInspectNextItemEmpty(t *testing.T) {
	myStack := Stack[int]{}
	_, err := myStack.InspectNextItem()
	if err == nil {
		t.Error("Error should not be nil, but", err)
	}
}
