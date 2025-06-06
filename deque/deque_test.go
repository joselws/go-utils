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

func TestAppendLeft(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendLeft(5)
	deque.AppendLeft(3)
	if deque.Len() != 2 {
		t.Errorf("Deque length should be 2, not %v", deque.Len())
	}
}

func TestAppendLeftFull(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendLeft(5)
	deque.AppendLeft(3)
	deque.AppendLeft(1)
	err := deque.AppendLeft(2)
	if err == nil {
		t.Errorf("AppendLeft should return %s", ErrFullDeque)
	}
}

func TestPopLeft(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendLeft(5)
	deque.AppendLeft(3)
	value, err := deque.PopLeft()
	if err != nil {
		t.Errorf("PopLeft should not return an error: %s", err)
	}
	if value != 3 {
		t.Errorf("PopLeft should return 3, not %v", value)
	}
	if deque.Len() != 1 {
		t.Errorf("Deque length should be 1, not %v", deque.Len())
	}
}

func TestPopLeftEmpty(t *testing.T) {
	deque := NewDeque[int](3)
	_, err := deque.PopLeft()
	if err == nil {
		t.Errorf("PopLeft should return %s", ErrEmptyDeque)
	}
}

func TestAppendLeftRight(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendLeft(5)
	deque.AppendRight(3)
	deque.AppendLeft(1)
	val1, err := deque.PopLeft()   // should be 1
	val2, err2 := deque.PopRight() // should be 3
	if deque.Len() != 1 {
		t.Errorf("Deque length should be 2, not %v", deque.Len())
	}
	if val1 != 1 {
		t.Errorf("PopLeft should return 1, not %v", val1)
	}
	if err != nil {
		t.Errorf("PopLeft should not return an error: %s", err)
	}
	if val2 != 3 {
		t.Errorf("PopRight should return 3, not %v", val2)
	}
	if err2 != nil {
		t.Errorf("PopRight should not return an error: %s", err2)
	}

}

func TestPeekRight(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	value, err := deque.PeekRight()
	if err != nil {
		t.Errorf("PeekRight should not return an error: %s", err)
	}
	if value != 3 {
		t.Errorf("PeekRight should return 3, not %v", value)
	}
	if deque.Len() != 2 {
		t.Errorf("Deque length should be 2, not %v", deque.Len())
	}
}

func TestPeekRightEmpty(t *testing.T) {
	deque := NewDeque[int](3)
	_, err := deque.PeekRight()
	if err == nil {
		t.Errorf("PeekRight should return %s", ErrEmptyDeque)
	}
}

func TestPeekLeft(t *testing.T) {
	deque := NewDeque[int](3)
	deque.AppendRight(5)
	deque.AppendRight(3)
	value, err := deque.PeekLeft()
	if err != nil {
		t.Errorf("PeekRight should not return an error: %s", err)
	}
	if value != 5 {
		t.Errorf("PeekRight should return 5, not %v", value)
	}
	if deque.Len() != 2 {
		t.Errorf("Deque length should be 2, not %v", deque.Len())
	}
}

func TestPeekLeftEmpty(t *testing.T) {
	deque := NewDeque[int](3)
	_, err := deque.PeekLeft()
	if err == nil {
		t.Errorf("PeekLeft should return %s", ErrEmptyDeque)
	}
}

func TestClearEmptiesDeque(t *testing.T) {
	deque := NewDeque[int](5)
	deque.AppendRight(1)
	deque.AppendLeft(2)
	deque.AppendRight(3)
	if deque.Len() != 3 {
		t.Errorf("Deque length should be 3 before Clear, got %v", deque.Len())
	}
	deque.Clear()
	if !deque.IsEmpty() {
		t.Error("Deque should be empty after Clear")
	}
	if deque.Len() != 0 {
		t.Errorf("Deque length should be 0 after Clear, got %v", deque.Len())
	}
	// Ensure operations work after Clear
	err := deque.AppendRight(4)
	if err != nil {
		t.Errorf("AppendRight after Clear should not return error: %v", err)
	}
	val, err := deque.PopLeft()
	if err != nil {
		t.Errorf("PopLeft after Clear should not return error: %v", err)
	}
	if val != 4 {
		t.Errorf("PopLeft after Clear should return 4, got %v", val)
	}
}
