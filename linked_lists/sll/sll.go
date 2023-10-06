/*
Singly linked list implementation.
Has only a pointer to Head and can only move to Next.
*/

package sll

import (
	"errors"
	"fmt"
)

// Each element on the linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// Create a new node by passing in a value.
func NewSingleLinkedListNode[T comparable](value T) *Node[T] {
	node := Node[T]{Value: value}
	return &node
}

// String representation of the Node for Stringer interface.
// Prints Node(value)(address-or-nil)
func (node *Node[T]) String() string {
	var nodeNext string
	if node.Next == nil {
		nodeNext = "nil"
	} else {
		nodeNext = "Node"
	}
	return fmt.Sprintf("Node(%v)(%v)", node.Value, nodeNext)
}

// Singly linked list type. Head can be a Node or nil.
// Length increases or decreases automatically.
type LinkedList[T comparable] struct {
	Head   *Node[T]
	Length int
}

// Creates a new singly linked list (by reference).
func NewSingleLinkedList[T comparable]() *LinkedList[T] {
	linkedList := LinkedList[T]{}
	return &linkedList
}

// String representation of the Linked List for Stringer interface.
// Prints "SinglyLinkedList(length)".
func (linkedList LinkedList[T]) String() string {
	return fmt.Sprintf("SinglyLinkedList(%v)", linkedList.Length)
}

// Add an element to the beginning of the Linked List.
// O(1) time complexity.
func (linkedList *LinkedList[T]) Prepend(value T) {
	newNode := NewSingleLinkedListNode(value)
	newNode.Next = linkedList.Head
	linkedList.Head = newNode
	linkedList.Length++
}

// Returns true if the value is within the Linked List.
// O(n) time complexity.
func (linkedList LinkedList[T]) Contains(value any) bool {
	for linkedList.Head != nil {
		if linkedList.Head.Value == value {
			return true
		}
		linkedList.Head = linkedList.Head.Next
	}
	return false
}

// Prints all the values of the elements in the Linked List.
// O(n) time complexity.
func (linkedList LinkedList[T]) PrintList() {
	for linkedList.Head != nil {
		fmt.Print(linkedList.Head.Value, " ")
		linkedList.Head = linkedList.Head.Next
	}
	fmt.Println()
}

// Deletes an element of the Linked List.
// Returns true if the value was found and deleted.
// Returns false if no value was found, thus none was deleted.
// O(n) time complexity.
func (linkedList *LinkedList[T]) Delete(value T) bool {

	if linkedList.Head == nil {
		return false
	}

	head := linkedList.Head
	var previous *Node[T]
	for linkedList.Head != nil {
		// delete first element
		if linkedList.Head.Value == value && previous == nil {
			linkedList.Head = linkedList.Head.Next
			linkedList.Length--
			return true
			// delete last element
		} else if linkedList.Head.Value == value && linkedList.Head.Next == nil {
			linkedList.Head = previous
			linkedList.Head.Next = nil
			linkedList.Head = head
			linkedList.Length--
			return true
			// delete middle elements
		} else if linkedList.Head.Value == value {
			previous.Next = linkedList.Head.Next
			linkedList.Head = head
			linkedList.Length--
			return true
			// check next node
		} else {
			previous = linkedList.Head
			linkedList.Head = linkedList.Head.Next
		}
	}
	linkedList.Head = head
	return false
}

// Removes and returns the current Head value of the list.
// Returns an error if the list is empty.
// O(1) time complexity.
func (linkedList *LinkedList[T]) PopFirst() (T, error) {
	var value T
	if linkedList.Head == nil {
		return value, errors.New("can't Pop from an empty linked list")
	}

	value = linkedList.Head.Value
	linkedList.Delete(value)
	return value, nil
}
