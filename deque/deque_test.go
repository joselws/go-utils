package deque

import "testing"

func TestIsFullNoSize(t *testing.T) {
	deque := NewDeque[int](0)
	isFull := deque.IsFull()
	if isFull {
		t.Error("isFull should be false on no size deques")
	}
}

func TestIsFullTrue(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	deque.AppendRight(1)
	if !deque.IsFull() {
		t.Error("isFull should be true")
	}
}

func TestIsFullFalse(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	if deque.IsFull() {
		t.Error("isFull should be false")
	}
}

func TestDequeEmpty(t *testing.T) {
	deque := NewDeque[int](3)
	if !deque.IsEmpty() {
		t.Error("Deque should be empty")
	}
}

func TestDequeNotEmpty(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	if deque.IsEmpty() {
		t.Error("Deque should not be empty")
	}
}

func TestAppendRight(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	if deque.Len() != 2 {
		t.Errorf("Deque length should be 2, not %v", deque.Len())
	}
}

func TestAppendRightFull(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	deque.AppendRight(1)
	err := deque.AppendRight(2)
	if err == nil {
		t.Errorf("AppendRight should return %s", ErrFullDeque)
	}
}

func TestPopRight(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	value, err := deque.PopRight()
	if err != nil {
		t.Errorf("PopRight should not return an error: %s", err)
	}
	if value != 3 {
		t.Errorf("PopRight should return 3, not %v", value)
	}
	if deque.Len() != 1 {
		t.Errorf("Deque length should be 1, not %v", deque.Len())
	}
}

func TestPopRightEmpty(t *testing.T) {
	deque := NewDeque[int](3)
	_, err := deque.PopRight()
	if err == nil {
		t.Errorf("PopRight should return %s", ErrEmptyDeque)
	}
}
