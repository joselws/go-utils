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
	isFull := deque.IsFull()
	if !isFull {
		t.Error("isFull should be true")
	}
}

func TestIsFullFalse(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	isFull := deque.IsFull()
	if isFull {
		t.Error("isFull should be false")
	}
}

func TestAppendRight(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	if deque.Len() != 2 {
		t.Error("")
	}
}

func TestAppendRightFull(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	deque.AppendRight(1)
	if deque.AppendRight(2) {
		t.Error("AppendRight should return false")
	}
}
