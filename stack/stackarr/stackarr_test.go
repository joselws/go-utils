package stackarr

import (
	"testing"
)

func TestStackCorrectlyInitialized(t *testing.T) {
	myStack := NewStack[int]()
	if myStack.Len() != 0 {
		t.Error("Stack length should be 0, not", myStack.Len())
	}
}

func TestPushMethod(t *testing.T) {
	myStack := NewStack[int]()
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	if myStack.Len() != 3 {
		t.Error("Stack length should be 3, not", myStack.Len())
	}
}

func TestPopMethod(t *testing.T) {
	myStack := NewStack[int]()
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	lastItem, err := myStack.Pop()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if myStack.Len() != 2 {
		t.Error("Stack length should be 2, not", myStack.Len())
	}
	if lastItem != 300 {
		t.Error("Last Stack item should be 300, not", lastItem)
	}
}

func TestPopOnEmptyStack(t *testing.T) {
	myStack := NewStack[int]()
	_, err := myStack.Pop()
	if err == nil {
		t.Error("Error should not be nil, but", err)
	}
}
