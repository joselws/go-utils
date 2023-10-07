/*
Doubly linked list implementation.
Has pointer to Head and Tail.
Can move forward (Next) and backwards (Prev).
*/

package dll

import (
	"errors"
	"fmt"
)

// Each element on the linked list.
type Node[T comparable] struct {
	Value T
	Next  *Node[T]
	Prev  *Node[T]
}

// Create a new node by passing in a value.
func NewDoublyLinkedListNode[T comparable](value T) *Node[T] {
	node := Node[T]{Value: value}
	return &node
}

// String representation of the Node for Stringer interface.
// Prints Node(value)(Next:Node|nil)(Prev:Node|nil)
func (node Node[T]) String() string {
	var nodeNext string
	if node.Next == nil {
		nodeNext = "nil"
	} else {
		nodeNext = "Node"
	}
	var nodePrev string
	if node.Prev == nil {
		nodePrev = "nil"
	} else {
		nodePrev = "Node"
	}
	return fmt.Sprintf("%v<-%v->%v", nodePrev, node.Value, nodeNext)
}

// Doubly linked list type. Head and Tail can be a Node or nil.
// Length increases or decreases automatically.
type LinkedList[T comparable] struct {
	Head   *Node[T]
	Tail   *Node[T]
	Length int
}

// Creates a new doubly linked list (by reference).
func NewDoublyLinkedList[T comparable]() *LinkedList[T] {
	linkedList := LinkedList[T]{}
	return &linkedList
}

// String representation of the Linked List for Stringer interface.
// Prints "DoublyLinkedList(length)".
func (linkedList LinkedList[T]) String() string {
	return fmt.Sprintf("DoublyLinkedList(%v)", linkedList.Length)
}

// Add an element at the beginning of the List.
// O(1) time complexity.
func (linkedList *LinkedList[T]) AppendLeft(value T) {
	newNode := NewDoublyLinkedListNode(value)
	if linkedList.Length == 0 {
		linkedList.Head = newNode
		linkedList.Tail = newNode
	} else {
		newNode.Next = linkedList.Head
		linkedList.Head.Prev = newNode
		linkedList.Head = newNode
	}
	linkedList.Length++
}

// Add an element at the end of the List.
// O(1) time complexity.
func (linkedList *LinkedList[T]) AppendRight(value T) {
	newNode := NewDoublyLinkedListNode(value)
	if linkedList.Length == 0 {
		linkedList.Head = newNode
		linkedList.Tail = newNode
	} else {
		newNode.Prev = linkedList.Tail
		linkedList.Tail.Next = newNode
		linkedList.Tail = newNode
	}
	linkedList.Length++
}

// Returns true if the value is within the Linked List.
// O(n) time complexity.
func (linkedList *LinkedList[T]) Contains(value any) bool {
	node := linkedList.Head
	for node != nil {
		if node.Value == value {
			return true
		}
		node = node.Next
	}
	return false
}

// Prints all the values of the elements in the Linked List.
// O(n) time complexity.
func (linkedList *LinkedList[T]) PrintList() {
	node := linkedList.Head
	for node != nil {
		fmt.Print(node.Value, " ")
		node = node.Next
	}
	fmt.Println()
}

// Deletes an element of the Linked List.
// Returns true if the value was found and deleted.
// Returns false if no value was found, thus none was deleted.
// O(n) time complexity.
func (linkedList *LinkedList[T]) Delete(value T) bool {

	// Empty list
	if linkedList.Head == nil {
		return false
	}

	// List with 1 element
	if linkedList.Length == 1 {
		linkedList.Head = nil
		linkedList.Tail = nil
		linkedList.Length--
		return true
	}

	// List with many elements
	// head := linkedList.Head
	node := linkedList.Head
	for node != nil {
		// delete first element
		if node.Value == value && node == linkedList.Head {
			linkedList.Head = linkedList.Head.Next
			linkedList.Head.Prev = nil
			linkedList.Length--
			return true
			// delete last element
		} else if node.Value == value && node == linkedList.Tail {
			linkedList.Tail = linkedList.Tail.Prev
			linkedList.Tail.Next = nil
			// linkedList.Head = head
			linkedList.Length--
			return true
			// delete middle elements
		} else if node.Value == value {
			node.Next.Prev = node.Prev
			node.Prev.Next = node.Next
			// linkedList.Head = head
			linkedList.Length--
			return true
			// check next node
		} else {
			node = node.Next
		}
	}
	// linkedList.Head = head
	return false
}

// Removes and returns the current Head value of the list.
// Returns an error if the list is empty.
// O(1) time complexity.
func (linkedList *LinkedList[T]) PopLeft() (T, error) {
	var value T
	if linkedList.Head == nil {
		return value, errors.New("can't Pop from an empty linked list")
	}

	value = linkedList.Head.Value
	// single element case
	if linkedList.Length == 1 {
		linkedList.Head = nil
		linkedList.Tail = nil
		// more than 1 element case
	} else {
		linkedList.Head = linkedList.Head.Next
		linkedList.Head.Prev = nil
	}
	linkedList.Length--
	return value, nil
}

// Removes and returns the current Tail value of the list.
// Returns an error if the list is empty.
// O(1) time complexity.
func (linkedList *LinkedList[T]) PopRight() (T, error) {
	var value T
	if linkedList.Tail == nil {
		return value, errors.New("can't Pop from an empty linked list")
	}

	value = linkedList.Tail.Value
	// single element case
	if linkedList.Length == 1 {
		linkedList.Head = nil
		linkedList.Tail = nil
		// more than 1 element case
	} else {
		linkedList.Tail = linkedList.Tail.Prev
		linkedList.Tail.Next = nil
	}
	linkedList.Length--
	return value, nil
}
