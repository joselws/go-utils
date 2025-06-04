package queue

import (
	"errors"
	"testing"
)

func TestNewQueue(t *testing.T) {
	q := NewQueue[int](10)
	if q == nil {
		t.Fatal("expected non-nil queue")
	}
	if q.Len() != 0 {
		t.Errorf("expected length 0, got %d", q.Len())
	}
}

func TestEnqueueDequeue(t *testing.T) {
	q := NewQueue[int](0) // unbounded
	for i := range 5 {
		if err := q.Enqueue(i); err != nil {
			t.Errorf("unexpected error on enqueue: %v", err)
		}
	}
	if q.Len() != 5 {
		t.Errorf("expected length 5, got %d", q.Len())
	}
	for i := range 5 {
		val, err := q.Dequeue()
		if err != nil {
			t.Errorf("unexpected error on dequeue: %v", err)
		}
		if val != i {
			t.Errorf("expected %d, got %d", i, val)
		}
	}
	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}
}

func TestEnqueueFullQueue(t *testing.T) {
	q := NewQueue[int](2)
	_ = q.Enqueue(1)
	_ = q.Enqueue(2)
	err := q.Enqueue(3)
	if !errors.Is(err, ErrFullQueue) {
		t.Errorf("expected ErrFullQueue, got %v", err)
	}
}

func TestDequeueEmptyQueue(t *testing.T) {
	q := NewQueue[int](2)
	_, err := q.Dequeue()
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("expected ErrEmptyQueue, got %v", err)
	}
}

func TestPeekNextAndPeekLast(t *testing.T) {
	q := NewQueue[int](0)
	_, err := q.PeekNext()
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("expected ErrEmptyQueue, got %v", err)
	}
	_, err = q.PeekLast()
	if !errors.Is(err, ErrEmptyQueue) {
		t.Errorf("expected ErrEmptyQueue, got %v", err)
	}
	_ = q.Enqueue(10)
	_ = q.Enqueue(20)
	next, err := q.PeekNext()
	if err != nil || next != 10 {
		t.Errorf("expected 10, got %d, err: %v", next, err)
	}
	last, err := q.PeekLast()
	if err != nil || last != 20 {
		t.Errorf("expected 20, got %d, err: %v", last, err)
	}
}

func TestIsFullAndIsEmpty(t *testing.T) {
	q := NewQueue[int](2)
	if !q.IsEmpty() {
		t.Error("expected queue to be empty")
	}
	_ = q.Enqueue(1)
	if q.IsEmpty() {
		t.Error("expected queue to be non-empty")
	}
	_ = q.Enqueue(2)
	if !q.IsFull() {
		t.Error("expected queue to be full")
	}
	_, _ = q.Dequeue()
	if q.IsFull() {
		t.Error("expected queue to not be full")
	}
}

func TestClearQueue(t *testing.T) {
	q := NewQueue[int](5)
	for i := range 5 {
		_ = q.Enqueue(i)
	}
	if q.Len() != 5 {
		t.Errorf("expected length 5, got %d", q.Len())
	}
	q.Clear()
	if !q.IsEmpty() {
		t.Error("expected queue to be empty after clear")
	}
	if q.Len() != 0 {
		t.Errorf("expected length 0, got %d", q.Len())
	}
}
