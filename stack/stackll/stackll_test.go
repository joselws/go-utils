package stackll

import (
	"fmt"
	"testing"
)

func TestCreateNewStack(t *testing.T) {
	stack := NewStack[int](20)
	if stack.size != 20 {
		t.Error("Stack size should be 20, not", stack.size)
	}
	if stack.linkedList.Head != nil {
		t.Error("Stack first element should be nil")
	}
	if stack.linkedList.Length != 0 {
		t.Error("Stack should have 0 elements, not", stack.linkedList.Length)
	}
}

func TestStringEmptyNoSize(t *testing.T) {
	stack := NewStack[int](0)
	stackString := fmt.Sprint(stack)
	if stackString != "Stack(0)" {
		t.Error("Stack string should be Stack(0), not", stackString)
	}
}

func TestStringNoSize(t *testing.T) {
	stack := NewStack[int](0)
	stack.Push(5)
	stack.Push(3)
	stack.Push(1)
	stackString := fmt.Sprint(stack)
	if stackString != "Stack(3)" {
		t.Error("Stack string should be Stack(3), not", stackString)
	}
}

func TestStringEmpty(t *testing.T) {
	stack := NewStack[int](3)
	stackString := fmt.Sprint(stack)
	if stackString != "Stack(0/3)" {
		t.Error("Stack string should be Stack(0/3), not", stackString)
	}
}

func TestStringSize(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	stack.Push(3)
	stackString := fmt.Sprint(stack)
	if stackString != "Stack(2/3)" {
		t.Error("Stack string should be Stack(2/3), not", stackString)
	}
}

func TestIsFullNoSize(t *testing.T) {
	stack := NewStack[int](0)
	isFull := stack.IsFull()
	if isFull {
		t.Error("isFull should be false on no size stacks")
	}
}

func TestIsFullTrue(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	stack.Push(3)
	stack.Push(1)
	isFull := stack.IsFull()
	if !isFull {
		t.Error("isFull should be true")
	}
}

func TestIsFullFalse(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	stack.Push(3)
	isFull := stack.IsFull()
	if isFull {
		t.Error("isFull should be false")
	}
}

func TestIsEmptyNoData(t *testing.T) {
	stack := NewStack[int](3)
	isEmpty := stack.IsEmpty()
	if !isEmpty {
		t.Error("isEmpty should be true")
	}
}

func TestIsEmptyWithData(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	isEmpty := stack.IsEmpty()
	if isEmpty {
		t.Error("isEmpty should be false")
	}
}

func TestPushNotFull(t *testing.T) {
	stack := NewStack[int](3)
	ok := stack.Push(5)
	if !ok {
		t.Error("Push should return true on not full stacks")
	}
	if stack.linkedList.Length != 1 {
		t.Error("Stack length should be 1, not", stack.linkedList.Length)
	}
}

func TestPushFull(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	stack.Push(3)
	stack.Push(1)
	ok := stack.Push(0)
	if ok {
		t.Error("Push should return false on full stacks")
	}
	if stack.linkedList.Length != 3 {
		t.Error("Stack length should be 3, not", stack.linkedList.Length)
	}
}

func TestLenData(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	if stack.Len() != 1 {
		t.Error("Stack length should be 1, not", stack.Len())
	}
}

func TestLenEmpty(t *testing.T) {
	stack := NewStack[int](3)
	if stack.Len() != 0 {
		t.Error("Stack length should be 0, not", stack.Len())
	}
}

func TestPopError(t *testing.T) {
	stack := NewStack[int](3)
	_, err := stack.Pop()
	if err == nil {
		t.Error("Error should be given, not nil")
	}
}

func TestPopValid(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	stack.Push(3)
	value, err := stack.Pop()
	if value != 3 {
		t.Error("Popped value should be 3, not", value)
	}
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if stack.Len() != 1 {
		t.Error("Stack length should be 1, not", stack.Len())
	}
}

func TestPeekError(t *testing.T) {
	stack := NewStack[int](3)
	_, err := stack.PeekNext()
	if err == nil {
		t.Error("Error should be given when peeking empty stack, not nil")
	}
}

func TestPeekValid(t *testing.T) {
	stack := NewStack[int](3)
	stack.Push(5)
	stack.Push(3)
	value, err := stack.PeekNext()
	if value != 3 {
		t.Error("Peeked value should be 3, not", value)
	}
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if stack.Len() != 2 {
		t.Error("Stack length should be 2, not", stack.Len())
	}
}
