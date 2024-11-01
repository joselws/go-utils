package stackarr

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

func TestStackCorrectlyCreated(t *testing.T) {
	myStack := NewStack[int](3)
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

func TestLen(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	var stackLen int = myStack.Len()
	if stackLen != 3 {
		t.Error("Stack length should be 3, not", stackLen)
	}
}

func TestString(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	var stackString string = fmt.Sprint(myStack)
	if stackString != "Stack[int][100 200 300]" {
		t.Error("Stack string representation should be Stack[int][100 200 300], not", stackString)
	}
}

func TestPeekNext(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	nextItem, err := myStack.PeekNext()
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

func TestPeekNextEmpty(t *testing.T) {
	myStack := Stack[int]{}
	_, err := myStack.PeekNext()
	if err == nil {
		t.Error("Error should not be nil, but", err)
	}
}

func TestExtractMany(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 300, 200}}
	items, err := myStack.ExtractMany(2)
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if len(items) != 2 {
		t.Error("Items extracted should be 2, not", items)
	}
	if items[0] != 200 {
		t.Error("First item of items should be 200, not", items[0])
	}
	if items[1] != 300 {
		t.Error("Second item of items should be 300, not", items[1])
	}
	if myStack.Len() != 1 {
		t.Error("Stack length should now be 1, not", myStack.Len())
	}
}

func TestExtractManyExceed(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	_, err := myStack.ExtractMany(4)
	if err == nil {
		t.Error("Error should not be nil, but", err)
	}
	if myStack.Len() != 3 {
		t.Error("Stack length should now be 3, not", myStack.Len())
	}
}

func TestInsertSlice(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	mySlice := []int{400, 500}
	myStack.InsertFromSlice(mySlice)
	if myStack.Len() != 5 {
		t.Error("Stack length should now be 5, not", myStack.Len())
	}
	item, _ := myStack.PeekNext()
	if item != 500 {
		t.Error("Last item in Stack should be 500, not", item)
	}
}

func TestInsertEmptySlice(t *testing.T) {
	myStack := Stack[int]{Items: []int{100, 200, 300}}
	mySlice := []int{}
	myStack.InsertFromSlice(mySlice)
	if myStack.Len() != 3 {
		t.Error("Stack length should now be 3, not", myStack.Len())
	}
	item, _ := myStack.PeekNext()
	if item != 300 {
		t.Error("Last item in Stack should be 300, not", item)
	}
}
