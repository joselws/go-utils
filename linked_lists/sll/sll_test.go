package sll

import (
	"fmt"
	"testing"
)

func TestNewSingleLinkedListNode(t *testing.T) {
	node := NewSingleLinkedListNode(10)
	if node.Value != 10 {
		t.Error("Node value should be 10, not", node.Value)
	}
	if node.Next != nil {
		t.Error("Node Next should be nil, not", node.Next)
	}
}

func TestNodeStringNil(t *testing.T) {
	node := NewSingleLinkedListNode(20)
	var nodeString string = fmt.Sprint(node)
	if fmt.Sprintf(nodeString) != "Node(20)(nil)" {
		t.Error("Node string repr should be Node(20)(nil), not", nodeString)
	}
}

func TestNodeStringWithNext(t *testing.T) {
	node := NewSingleLinkedListNode(20)
	node2 := NewSingleLinkedListNode(30)
	node.Next = node2
	var nodeString string = fmt.Sprint(node)
	if fmt.Sprintf(nodeString) != "Node(20)(Node)" {
		t.Error("Node string repr should be Node(20)(Node), not", nodeString)
	}
}

func TestNewSingleLinkedList(t *testing.T) {
	linkedList := NewSingleLinkedList()
	if linkedList.Length != 0 {
		t.Error("linkedList Length should be 0, not", linkedList.Length)
	}
	if linkedList.Head != nil {
		t.Error("linkedList Head should be nil, not", linkedList.Head)
	}
}

func TestLinkedListString(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedListString := fmt.Sprint(linkedList)
	if linkedListString != "SinglyLinkedList(0)" {
		t.Error("linkedList string should be SinglyLinkedList(0), not", linkedListString)
	}
}

func TestLinkedListPrepend(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	if linkedList.Length != 1 {
		t.Error("Linked List length should be 1, not", linkedList.Length)
	}
	if linkedList.Head.Value != 20 {
		t.Error("Linked List head value should be 20, not", linkedList.Head.Value)
	}
	if linkedList.Head.Next != nil {
		t.Error("Linked List head next should be nil, not", linkedList.Head.Next)
	}
}

func TestLinkedListPrependMany(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	linkedList.Prepend(10)
	if linkedList.Length != 2 {
		t.Error("Linked List length should be 2, not", linkedList.Length)
	}
	if linkedList.Head.Value != 10 {
		t.Error("Linked List head value should be 10, not", linkedList.Head.Value)
	}
	if linkedList.Head.Next == nil {
		t.Error("Linked List head next should not be nil, but", linkedList.Head.Next)
	}
	if linkedList.Head.Next.Value != 20 {
		t.Error("Linked List head value should be 20, not", linkedList.Head.Next.Value)
	}
	if linkedList.Head.Next.Next != nil {
		t.Error("Linked List head next should be nil, not", linkedList.Head.Next.Next)
	}
}

func TestEmptyLinkedListContains(t *testing.T) {
	linkedList := NewSingleLinkedList()
	if linkedList.Contains(20) {
		t.Error("Empty Linked list should not contain 20")
	}
}

func TestLinkedListContains(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	linkedList.Prepend(10)
	if !linkedList.Contains(20) {
		t.Error("Empty Linked list should contain 20")
	}
	if linkedList.Contains(50) {
		t.Error("Empty Linked list should not contain 50")
	}
}

func TestDeleteEmptyLinkedList(t *testing.T) {
	linkedList := NewSingleLinkedList()
	if linkedList.Delete(20) {
		t.Error("Linked list should not delete 20 if empty")
	}
	if linkedList.Length != 0 {
		t.Error("Linked list length should be 0, not", linkedList.Length)
	}
}

func TestDeleteFirstElement(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	linkedList.Prepend(10)
	linkedList.Prepend(50)
	if !linkedList.Delete(50) {
		t.Error("Linked list should have deleted 50")
	}
	if linkedList.Length != 2 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 10 {
		t.Error("Linked list head value should now be 10, not", linkedList.Head.Value)
	}
}

func TestDeleteLastElement(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	linkedList.Prepend(10)
	linkedList.Prepend(50)
	if !linkedList.Delete(20) {
		t.Error("Linked list should have deleted 20")
	}
	if linkedList.Length != 2 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 50 {
		t.Error("Linked list head value should still be 50, not", linkedList.Head.Value)
	}
	if linkedList.Head.Next.Value != 10 {
		t.Error("Linked list next head value should still be 10, not", linkedList.Head.Value)
	}
}

func TestDeleteMiddleElement(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	linkedList.Prepend(10)
	linkedList.Prepend(50)
	if !linkedList.Delete(10) {
		t.Error("Linked list should have deleted 10")
	}
	if linkedList.Length != 2 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 50 {
		t.Error("Linked list head value should still be 50, not", linkedList.Head.Value)
	}
	if linkedList.Head.Next.Value != 20 {
		t.Error("Linked list next head value should still be 20, not", linkedList.Head.Value)
	}
}

func TestDeleteElementNotFound(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	linkedList.Prepend(10)
	linkedList.Prepend(50)
	if linkedList.Delete(9) {
		t.Error("Linked list should not have deleted 9")
	}
	if linkedList.Length != 3 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 50 {
		t.Error("Linked list head value should still be 50, not", linkedList.Head.Value)
	}
}

func TestPopFirst(t *testing.T) {
	linkedList := NewSingleLinkedList()
	linkedList.Prepend(20)
	value, err := linkedList.PopFirst()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if value != 20 {
		t.Error("Popped value should be 20, not", value)
	}
	if linkedList.Length != 0 {
		t.Error("Linked List length should be 0, not", linkedList.Length)
	}
}

func TestPopFirstError(t *testing.T) {
	linkedList := NewSingleLinkedList()
	_, err := linkedList.PopFirst()
	if err == nil {
		t.Error("Error should not be nil")
	}
	if linkedList.Length != 0 {
		t.Error("Linked List length should be 0, not", linkedList.Length)
	}
}
