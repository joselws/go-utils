package sll

import (
	"errors"
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func NewSingleLinkedListNode(value interface{}) *Node {
	node := Node{Value: value}
	return &node
}

func (node *Node) String() string {
	var nodeNext string
	if node.Next == nil {
		nodeNext = "nil"
	} else {
		nodeNext = "Node"
	}
	return fmt.Sprintf("Node(%v)(%v)", node.Value, nodeNext)
}

type LinkedList struct {
	Head   *Node
	Length uint
}

func NewSingleLinkedList() *LinkedList {
	linkedList := LinkedList{}
	return &linkedList
}

func (linkedList LinkedList) String() string {
	return fmt.Sprintf("SinglyLinkedList(%v)", linkedList.Length)
}

func (linkedList *LinkedList) Prepend(value interface{}) {
	newNode := NewSingleLinkedListNode(value)
	newNode.Next = linkedList.Head
	linkedList.Head = newNode
	linkedList.Length++
}

func (linkedList LinkedList) Contains(value interface{}) bool {
	for linkedList.Head != nil {
		if linkedList.Head.Value == value {
			return true
		}
		linkedList.Head = linkedList.Head.Next
	}
	return false
}

func (linkedList LinkedList) PrintList() {
	for linkedList.Head != nil {
		fmt.Print(linkedList.Head.Value, " ")
		linkedList.Head = linkedList.Head.Next
	}
	fmt.Println()
}

func (linkedList *LinkedList) Delete(value interface{}) bool {

	if linkedList.Head == nil {
		return false
	}

	head := linkedList.Head
	var previous *Node
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

func (linkedList *LinkedList) PopFirst() (interface{}, error) {
	var value interface{}
	if linkedList.Head == nil {
		return value, errors.New("can't Pop from an empty linked list")
	}

	value = linkedList.Head.Value
	linkedList.Delete(value)
	return value, nil
}
