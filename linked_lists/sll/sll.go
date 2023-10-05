package sll

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func NewSingleLinkedListNode[T comparable](value T) *Node[T] {
	node := Node[T]{Value: value}
	return &node
}

func (node *Node[T]) String() string {
	var nodeNext string
	if node.Next == nil {
		nodeNext = "nil"
	} else {
		nodeNext = "Node"
	}
	return fmt.Sprintf("Node(%v)(%v)", node.Value, nodeNext)
}

type LinkedList[T comparable] struct {
	Head   *Node[T]
	Length int
}

func NewSingleLinkedList[T comparable]() *LinkedList[T] {
	linkedList := LinkedList[T]{}
	return &linkedList
}

func (linkedList LinkedList[T]) String() string {
	return fmt.Sprintf("SinglyLinkedList(%v)", linkedList.Length)
}

func (linkedList *LinkedList[T]) Prepend(value T) {
	newNode := NewSingleLinkedListNode(value)
	newNode.Next = linkedList.Head
	linkedList.Head = newNode
	linkedList.Length++
}

func (linkedList LinkedList[T]) Contains(value any) bool {
	_, ok := value.(T)
	if ok {

	}

	for linkedList.Head != nil {
		if linkedList.Head.Value == value {
			return true
		}
		linkedList.Head = linkedList.Head.Next
	}
	return false
}

func (linkedList LinkedList[T]) PrintList() {
	for linkedList.Head != nil {
		fmt.Print(linkedList.Head.Value, " ")
		linkedList.Head = linkedList.Head.Next
	}
	fmt.Println()
}

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

func (linkedList *LinkedList[T]) PopFirst() (T, error) {
	var value T
	if linkedList.Head == nil {
		return value, errors.New("can't Pop from an empty linked list")
	}

	value = linkedList.Head.Value
	linkedList.Delete(value)
	return value, nil
}
