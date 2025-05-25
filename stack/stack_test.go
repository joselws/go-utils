package stack

import (
	"testing"
)

func TestNewStack(t *testing.T) {
	s := NewStack[int](5)
	if s == nil {
		t.Error("NewStack returned nil")
	}
	if s.Len() != 0 {
		t.Errorf("NewStack expected length 0, got %d", s.Len())
	}
}

func TestLen(t *testing.T) {
	s := NewStack[int](5)
	if s.Len() != 0 {
		t.Errorf("Expected initial length 0, got %d", s.Len())
	}

	s.Push(10)
	if s.Len() != 1 {
		t.Errorf("Expected length 1 after one push, got %d", s.Len())
	}

	s.Push(20)
	if s.Len() != 2 {
		t.Errorf("Expected length 2 after two pushes, got %d", s.Len())
	}

	s.Pop()
	if s.Len() != 1 {
		t.Errorf("Expected length 1 after one pop, got %d", s.Len())
	}
}

func TestPush(t *testing.T) {
	s := NewStack[int](2) // Stack with capacity 2

	// Test successful pushes
	err := s.Push(1)
	if err != nil {
		t.Errorf("Push(1) returned an unexpected error: %v", err)
	}
	if s.Len() != 1 {
		t.Errorf("Expected length 1 after push, got %d", s.Len())
	}

	err = s.Push(2)
	if err != nil {
		t.Errorf("Push(2) returned an unexpected error: %v", err)
	}
	if s.Len() != 2 {
		t.Errorf("Expected length 2 after push, got %d", s.Len())
	}

	// Test pushing to a full stack
	err = s.Push(3)
	if err != ErrFullStack {
		t.Errorf("Push(3) on full stack expected ErrFullStack, got %v", err)
	}
	if s.Len() != 2 { // Length should not change
		t.Errorf("Expected length 2 after failed push, got %d", s.Len())
	}
}

func TestPop(t *testing.T) {
	s := NewStack[int](3)

	// Test Pop on an empty stack
	_, err := s.Pop()
	if err != ErrEmptyStack {
		t.Errorf("Pop on empty stack expected ErrEmptyStack, got %v", err)
	}

	// Push some elements
	s.Push(10)
	s.Push(20)
	s.Push(30)

	// Test successful pops
	val, err := s.Pop()
	if err != nil {
		t.Errorf("Pop returned an unexpected error: %v", err)
	}
	if val != 30 {
		t.Errorf("Expected Pop to return 30, got %d", val)
	}
	if s.Len() != 2 {
		t.Errorf("Expected length 2 after pop, got %d", s.Len())
	}

	val, err = s.Pop()
	if err != nil {
		t.Errorf("Pop returned an unexpected error: %v", err)
	}
	if val != 20 {
		t.Errorf("Expected Pop to return 20, got %d", val)
	}
	if s.Len() != 1 {
		t.Errorf("Expected length 1 after pop, got %d", s.Len())
	}

	val, err = s.Pop()
	if err != nil {
		t.Errorf("Pop returned an unexpected error: %v", err)
	}
	if val != 10 {
		t.Errorf("Expected Pop to return 10, got %d", val)
	}
	if s.Len() != 0 {
		t.Errorf("Expected length 0 after pop, got %d", s.Len())
	}

	// Test Pop again on empty stack
	_, err = s.Pop()
	if err != ErrEmptyStack {
		t.Errorf("Pop on empty stack expected ErrEmptyStack, got %v", err)
	}
}

func TestPeekNext(t *testing.T) {
	s := NewStack[string](3)

	// Test PeekNext on an empty stack
	_, err := s.PeekNext()
	if err != ErrEmptyStack {
		t.Errorf("PeekNext on empty stack expected ErrEmptyStack, got %v", err)
	}

	s.Push("apple")
	s.Push("banana")

	// Test successful PeekNext
	val, err := s.PeekNext()
	if err != nil {
		t.Errorf("PeekNext returned an unexpected error: %v", err)
	}
	if val != "banana" {
		t.Errorf("Expected PeekNext to return 'banana', got '%s'", val)
	}
	if s.Len() != 2 { // Length should not change after PeekNext
		t.Errorf("Length changed after PeekNext, expected 2, got %d", s.Len())
	}

	s.Push("cherry")
	val, err = s.PeekNext()
	if err != nil {
		t.Errorf("PeekNext returned an unexpected error: %v", err)
	}
	if val != "cherry" {
		t.Errorf("Expected PeekNext to return 'cherry', got '%s'", val)
	}
}

func TestIsEmpty(t *testing.T) {
	s := NewStack[int](2)
	if !s.IsEmpty() {
		t.Errorf("Expected IsEmpty to be true on new stack, got false")
	}

	s.Push(1)
	if s.IsEmpty() {
		t.Errorf("Expected IsEmpty to be false after push, got true")
	}

	s.Pop()
	if !s.IsEmpty() {
		t.Errorf("Expected IsEmpty to be true after popping last element, got false")
	}
}

func TestIsFull(t *testing.T) {
	s := NewStack[int](2)
	if s.IsFull() {
		t.Errorf("Expected IsFull to be false on new stack, got true")
	}

	s.Push(1)
	if s.IsFull() {
		t.Errorf("Expected IsFull to be false after one push, got true")
	}

	s.Push(2)
	if !s.IsFull() {
		t.Errorf("Expected IsFull to be true after filling stack, got false")
	}

	s.Pop()
	if s.IsFull() {
		t.Errorf("Expected IsFull to be false after pop from full stack, got true")
	}
}
