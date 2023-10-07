package dll

import (
	"fmt"
	"testing"
)

func TestNewDoublyLinkedListNode(t *testing.T) {
	node := NewDoublyLinkedListNode[int](10)
	if node.Value != 10 {
		t.Error("Node value should be 10, not", node.Value)
	}
	if node.Next != nil {
		t.Error("Node Next should be nil, not", node.Next)
	}
	if node.Prev != nil {
		t.Error("Node Prev should be nil, not", node.Prev)
	}
}

func TestNodeStringNil(t *testing.T) {
	node := NewDoublyLinkedListNode[int](20)
	var nodeString string = fmt.Sprint(node)
	if nodeString != "nil<-20->nil" {
		t.Error("Node string repr should be nil<-20->nil, not", nodeString)
	}
}

func TestNodeStringWithNextAndPrev(t *testing.T) {
	node := NewDoublyLinkedListNode[int](20)
	node2 := NewDoublyLinkedListNode[int](30)
	node3 := NewDoublyLinkedListNode[int](40)
	node.Next = node2
	node.Prev = node3
	var nodeString string = fmt.Sprint(node)
	if nodeString != "Node<-20->Node" {
		t.Error("Node string repr should be Node<-20->Node, not", nodeString)
	}
}

func TestNewDoublyLinkedList(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	if linkedList.Length != 0 {
		t.Error("linkedList Length should be 0, not", linkedList.Length)
	}
	if linkedList.Head != nil {
		t.Error("linkedList Head should be nil, not", linkedList.Head)
	}
	if linkedList.Tail != nil {
		t.Error("linkedList Tail should be nil, not", linkedList.Head)
	}
}

func TestLinkedListString(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedListString := fmt.Sprint(linkedList)
	if linkedListString != "DoublyLinkedList(0)" {
		t.Error("linkedList string should be DoublyLinkedList(0), not", linkedListString)
	}
}

func TestLinkedListAppendLeftSingleElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	if linkedList.Length != 1 {
		t.Error("Linked List length should be 1, not", linkedList.Length)
	}
	if linkedList.Head.Value != 20 {
		t.Error("Linked List head value should be 20, not", linkedList.Head.Value)
	}
	if linkedList.Head.Next != nil {
		t.Error("Linked List head next should be nil, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("Linked List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Value != 20 {
		t.Error("Linked List tail value should be 20, not", linkedList.Tail.Value)
	}
	if linkedList.Tail.Next != nil {
		t.Error("Linked List Tail next should be nil, not", linkedList.Tail.Next)
	}
	if linkedList.Tail.Prev != nil {
		t.Error("Linked List Tail prev should be nil, not", linkedList.Tail.Prev)
	}
}

func TestLinkedListAppendLeftMany(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	linkedList.AppendLeft(10)
	if linkedList.Length != 2 {
		t.Error("Linked List length should be 2, not", linkedList.Length)
	}
	if linkedList.Head.Value != 10 {
		t.Error("Linked List head value should be 10, not", linkedList.Head.Value)
	}
	if linkedList.Head != linkedList.Tail.Prev {
		t.Error("Linked List head should be tail prev, not", linkedList.Tail.Prev)
	}
	if linkedList.Tail != linkedList.Head.Next {
		t.Error("Linked List tail should be head next, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("Linked List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Value != 20 {
		t.Error("Linked List Tail value should be 20, not", linkedList.Tail.Value)
	}
	if linkedList.Tail.Next != nil {
		t.Error("Linked List Tail next should be nil, not", linkedList.Tail.Next)
	}
}

func TestLinkedListAppendRightSingleElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendRight(20)
	if linkedList.Length != 1 {
		t.Error("Linked List length should be 1, not", linkedList.Length)
	}
	if linkedList.Head.Value != 20 {
		t.Error("Linked List head value should be 20, not", linkedList.Head.Value)
	}
	if linkedList.Head.Next != nil {
		t.Error("Linked List head next should be nil, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("Linked List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Value != 20 {
		t.Error("Linked List tail value should be 20, not", linkedList.Tail.Value)
	}
	if linkedList.Tail.Next != nil {
		t.Error("Linked List Tail next should be nil, not", linkedList.Tail.Next)
	}
	if linkedList.Tail.Prev != nil {
		t.Error("Linked List Tail prev should be nil, not", linkedList.Tail.Prev)
	}
}

func TestLinkedListAppendRightMany(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendRight(20)
	linkedList.AppendRight(10)
	if linkedList.Length != 2 {
		t.Error("Linked List length should be 2, not", linkedList.Length)
	}
	if linkedList.Head.Value != 20 {
		t.Error("Linked List head value should be 20, not", linkedList.Head.Value)
	}
	if linkedList.Head != linkedList.Tail.Prev {
		t.Error("Linked List head should be tail prev, not", linkedList.Tail.Prev)
	}
	if linkedList.Tail != linkedList.Head.Next {
		t.Error("Linked List tail should be head next, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("Linked List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Value != 10 {
		t.Error("Linked List Tail value should be 10, not", linkedList.Tail.Value)
	}
	if linkedList.Tail.Next != nil {
		t.Error("Linked List Tail next should be nil, not", linkedList.Tail.Next)
	}
}

func TestEmptyLinkedListContains(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	if linkedList.Contains(20) {
		t.Error("Empty Linked list should not contain 20")
	}
}

func TestLinkedListContains(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendRight(20)
	linkedList.AppendRight(10)
	if !linkedList.Contains(20) {
		t.Error("Empty Linked list should contain 20")
	}
	if linkedList.Contains(50) {
		t.Error("Empty Linked list should not contain 50")
	}
}

func TestDeleteEmptyLinkedList(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	if linkedList.Delete(20) {
		t.Error("Linked list should not delete 20 if empty")
	}
	if linkedList.Length != 0 {
		t.Error("Linked list length should be 0, not", linkedList.Length)
	}
}

func TestDeleteFirstElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	linkedList.AppendLeft(10)
	linkedList.AppendLeft(50)
	if !linkedList.Delete(50) {
		t.Error("Linked list should have deleted 50")
	}
	if linkedList.Length != 2 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 10 {
		t.Error("Linked list head value should now be 10, not", linkedList.Head.Value)
	}
	if linkedList.Tail.Value != 20 {
		t.Error("Linked list tail value should now be 20, not", linkedList.Tail.Value)
	}
	if linkedList.Head != linkedList.Tail.Prev {
		t.Error("Linked List head should be tail prev, not", linkedList.Tail.Prev)
	}
	if linkedList.Tail != linkedList.Head.Next {
		t.Error("Linked List tail should be head next, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Next != nil {
		t.Error("List tail next should be nil, not", linkedList.Tail.Next)
	}
}

func TestDeleteLastElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	linkedList.AppendLeft(10)
	linkedList.AppendLeft(50)
	if !linkedList.Delete(20) {
		t.Error("Linked list should have deleted 20")
	}
	if linkedList.Length != 2 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 50 {
		t.Error("Linked list head value should still be 50, not", linkedList.Head.Value)
	}
	if linkedList.Tail.Value != 10 {
		t.Error("Linked list head value should still be 10, not", linkedList.Tail.Value)
	}
	if linkedList.Head != linkedList.Tail.Prev {
		t.Error("Linked List head should be tail prev, not", linkedList.Tail.Prev)
	}
	if linkedList.Tail != linkedList.Head.Next {
		t.Error("Linked List tail should be head next, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Next != nil {
		t.Error("List tail next should be nil, not", linkedList.Tail.Next)
	}
}

func TestDeleteMiddleElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	linkedList.AppendLeft(10)
	linkedList.AppendLeft(50)
	if !linkedList.Delete(10) {
		t.Error("Linked list should have deleted 10")
	}
	if linkedList.Length != 2 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 50 {
		t.Error("Linked list head value should still be 50, not", linkedList.Head.Value)
	}
	if linkedList.Tail.Value != 20 {
		t.Error("Linked list head value should still be 20, not", linkedList.Tail.Value)
	}
	if linkedList.Head != linkedList.Tail.Prev {
		t.Error("Linked List head should be tail prev, not", linkedList.Tail.Prev)
	}
	if linkedList.Tail != linkedList.Head.Next {
		t.Error("Linked List tail should be head next, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Next != nil {
		t.Error("List tail next should be nil, not", linkedList.Tail.Next)
	}
}

func TestDeleteElementNotFound(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	linkedList.AppendLeft(10)
	linkedList.AppendLeft(50)
	if linkedList.Delete(9) {
		t.Error("Linked list should not have deleted 9")
	}
	if linkedList.Length != 3 {
		t.Error("Linked list length should be 2")
	}
	if linkedList.Head.Value != 50 {
		t.Error("Linked list head value should still be 50, not", linkedList.Head.Value)
	}
	if linkedList.Tail.Value != 20 {
		t.Error("Linked list head value should still be 20, not", linkedList.Tail.Value)
	}
	if linkedList.Tail.Prev != linkedList.Head.Next {
		t.Error("Linked list head next and tail prev should now be equal")
	}
	if linkedList.Head.Next.Value != 10 {
		t.Error("Linked list head next value should still be 10, not", linkedList.Head.Next.Value)
	}
	if linkedList.Head.Prev != nil {
		t.Error("List head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Next != nil {
		t.Error("List tail next should be nil, not", linkedList.Tail.Next)
	}
}

func TestDeleteLinkedListSingleElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	if !linkedList.Delete(20) {
		t.Error("Linked list should have deleted 20")
	}
	if linkedList.Length != 0 {
		t.Error("Linked list length should be 0, not", linkedList.Length)
	}
	if linkedList.Head != nil {
		t.Error("Linked list head should now be nil, not", linkedList.Head)
	}
	if linkedList.Tail != nil {
		t.Error("Linked list Tail should now be nil, not", linkedList.Tail)
	}
}

func TestPopLeftManyElements(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	linkedList.AppendLeft(10)
	linkedList.AppendLeft(50)
	value, err := linkedList.PopLeft()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if value != 50 {
		t.Error("Popped value should be 50, not", value)
	}
	if linkedList.Length != 2 {
		t.Error("Linked List length should be 2, not", linkedList.Length)
	}
	if linkedList.Head.Value != 10 {
		t.Error("List head value should be 10, not", linkedList.Head.Value)
	}
	if linkedList.Tail.Value != 20 {
		t.Error("List Tail value should be 20, not", linkedList.Tail.Value)
	}
	if linkedList.Head != linkedList.Tail.Prev {
		t.Error("Linked List head should be tail prev, not", linkedList.Tail.Prev)
	}
	if linkedList.Tail != linkedList.Head.Next {
		t.Error("Linked List tail should be head next, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("list head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Next != nil {
		t.Error("list Tail Next should be nil, not", linkedList.Tail.Next)
	}
}

func TestPopLeftSingleElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	value, err := linkedList.PopLeft()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if value != 20 {
		t.Error("Popped value should be 20, not", value)
	}
	if linkedList.Length != 0 {
		t.Error("Linked List length should be 0, not", linkedList.Length)
	}
	if linkedList.Head != nil {
		t.Error("List head should be nil, not", linkedList.Head)
	}
	if linkedList.Tail != nil {
		t.Error("List Tail should be nil, not", linkedList.Tail)
	}
}

func TestPopLeftError(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	_, err := linkedList.PopLeft()
	if err == nil {
		t.Error("Error should be nil, not", err)
	}
	if linkedList.Length != 0 {
		t.Error("Linked List length should be 0, not", linkedList.Length)
	}
	if linkedList.Head != nil {
		t.Error("List head should be nil, not", linkedList.Head)
	}
	if linkedList.Tail != nil {
		t.Error("List Tail should be nil, not", linkedList.Tail)
	}
}

func TestPopRightManyElements(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	linkedList.AppendLeft(10)
	linkedList.AppendLeft(50)
	value, err := linkedList.PopRight()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if value != 20 {
		t.Error("Popped value should be 20, not", value)
	}
	if linkedList.Length != 2 {
		t.Error("Linked List length should be 2, not", linkedList.Length)
	}
	if linkedList.Head.Value != 50 {
		t.Error("List head value should be 50, not", linkedList.Head.Value)
	}
	if linkedList.Tail.Value != 10 {
		t.Error("List Tail value should be 10, not", linkedList.Tail.Value)
	}
	if linkedList.Head != linkedList.Tail.Prev {
		t.Error("Linked List head should be tail prev, not", linkedList.Tail.Prev)
	}
	if linkedList.Tail != linkedList.Head.Next {
		t.Error("Linked List tail should be head next, not", linkedList.Head.Next)
	}
	if linkedList.Head.Prev != nil {
		t.Error("list head prev should be nil, not", linkedList.Head.Prev)
	}
	if linkedList.Tail.Next != nil {
		t.Error("list Tail Next should be nil, not", linkedList.Tail.Next)
	}
}

func TestPopRightSingleElement(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	linkedList.AppendLeft(20)
	value, err := linkedList.PopRight()
	if err != nil {
		t.Error("Error should be nil, not", err)
	}
	if value != 20 {
		t.Error("Popped value should be 20, not", value)
	}
	if linkedList.Length != 0 {
		t.Error("Linked List length should be 0, not", linkedList.Length)
	}
	if linkedList.Head != nil {
		t.Error("List head should be nil, not", linkedList.Head)
	}
	if linkedList.Tail != nil {
		t.Error("List Tail should be nil, not", linkedList.Tail)
	}
}

func TestPopRightError(t *testing.T) {
	linkedList := NewDoublyLinkedList[int]()
	_, err := linkedList.PopRight()
	if err == nil {
		t.Error("Error should be nil, not", err)
	}
	if linkedList.Length != 0 {
		t.Error("Linked List length should be 0, not", linkedList.Length)
	}
	if linkedList.Head != nil {
		t.Error("List head should be nil, not", linkedList.Head)
	}
	if linkedList.Tail != nil {
		t.Error("List Tail should be nil, not", linkedList.Tail)
	}
}
