package queuell

import (
	"fmt"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := NewQueue[int](3)
	if queue.list.Length != 0 {
		t.Error("Initial queue length should be 0, not", queue.list.Length)
	}
	if queue.size != 3 {
		t.Error("Queue size should be 3, not", queue.size)
	}
}

func TestEnqueue(t *testing.T) {
	queue := NewQueue[int](3)
	if !queue.Enqueue(5) {
		t.Error("You should still enqueue in not full queues")
	}
	if !queue.Enqueue(10) {
		t.Error("You should still enqueue in not full queues")
	}
	if !queue.Enqueue(2) {
		t.Error("You should still enqueue in not full queues")
	}
	if queue.list.Head.Value != 5 {
		t.Error("Head value should be 5, not", queue.list.Head.Value)
	}
	if queue.list.Tail.Value != 2 {
		t.Error("Head value should be 2, not", queue.list.Tail.Value)
	}
	if queue.list.Length != 3 {
		t.Error("queue length should be 3, not", queue.list.Length)
	}
}

func TestFullEnqueue(t *testing.T) {
	queue := NewQueue[int](1)
	queue.Enqueue(5)
	if queue.Enqueue(3) {
		t.Error("You should not enqueue on full queues")
	}
	if queue.list.Length != 1 {
		t.Error("queue length should be 1, not", queue.list.Length)
	}
}

func TestEnqueueNoSize(t *testing.T) {
	queue := NewQueue[int](0)
	if !queue.Enqueue(3) {
		t.Error("You should enqueue on no size queues")
	}
	if queue.list.Length != 1 {
		t.Error("queue length should be 1, not", queue.list.Length)
	}
}

func TestIsFullNoSize(t *testing.T) {
	queue := NewQueue[int](0)
	isFull := queue.IsFull()
	if isFull {
		t.Error("isFull should be false on no size queues")
	}
}

func TestIsFullTrue(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(1)
	isFull := queue.IsFull()
	if !isFull {
		t.Error("isFull should be true")
	}
}

func TestIsFullFalse(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	queue.Enqueue(3)
	isFull := queue.IsFull()
	if isFull {
		t.Error("isFull should be false")
	}
}

func TestIsEmptyNoData(t *testing.T) {
	queue := NewQueue[int](3)
	isEmpty := queue.IsEmpty()
	if !isEmpty {
		t.Error("isEmpty should be true")
	}
}

func TestIsEmptyWithData(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	isEmpty := queue.IsEmpty()
	if isEmpty {
		t.Error("isEmpty should be false")
	}
}

func TestLenData(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	if queue.Len() != 1 {
		t.Error("Queue length should be 1, not", queue.Len())
	}
}

func TestLenEmpty(t *testing.T) {
	queue := NewQueue[int](3)
	if queue.Len() != 0 {
		t.Error("Queue length should be 0, not", queue.Len())
	}
}

func TestDequeueError(t *testing.T) {
	queue := NewQueue[int](3)
	_, ok := queue.Dequeue()
	if ok {
		t.Error("Queue should not dequeue if empty")
	}
	if queue.Len() != 0 {
		t.Error("Queue length should still be 0, not", queue.Len())
	}
}

func TestDequeue(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(1)

	value, ok := queue.Dequeue()
	if value != 5 {
		t.Error("Next value should be 5, not", value)
	}
	if !ok {
		t.Error("Ok should be true")
	}
	if queue.Len() != 2 {
		t.Error("Queue len should now be 2, not", queue.Len())
	}

	value, ok = queue.Dequeue()
	if value != 3 {
		t.Error("Next value should be 3, not", value)
	}
	if !ok {
		t.Error("Ok should be true")
	}
	if queue.Len() != 1 {
		t.Error("Queue len should now be 1, not", queue.Len())
	}
	value, ok = queue.Dequeue()
	if value != 1 {
		t.Error("Next value should be 1, not", value)
	}
	if !ok {
		t.Error("Ok should be true")
	}
	if queue.Len() != 0 {
		t.Error("Queue len should now be 2, not", queue.Len())
	}
}

func TestStringEmptyNoSize(t *testing.T) {
	queue := NewQueue[int](0)
	queueString := fmt.Sprint(queue)
	if queueString != "Queue(0)" {
		t.Error("Queue string should be Queue(0), not", queueString)
	}
}

func TestStringNoSize(t *testing.T) {
	queue := NewQueue[int](0)
	queue.Enqueue(5)
	queue.Enqueue(3)
	queue.Enqueue(1)
	queueString := fmt.Sprint(queue)
	if queueString != "Queue(3)" {
		t.Error("Queue string should be Queue(3), not", queueString)
	}
}

func TestStringEmpty(t *testing.T) {
	queue := NewQueue[int](3)
	queueString := fmt.Sprint(queue)
	if queueString != "Queue(0/3)" {
		t.Error("Queue string should be Queue(0/3), not", queueString)
	}
}

func TestStringSize(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	queue.Enqueue(3)
	queueString := fmt.Sprint(queue)
	if queueString != "Queue(2/3)" {
		t.Error("Queue string should be Queue(2/3), not", queueString)
	}
}

func TestPeekNextError(t *testing.T) {
	queue := NewQueue[int](3)
	_, ok := queue.PeekNext()
	if ok {
		t.Error("You can't peek on empty queue")
	}
}

func TestPeekNextValid(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	queue.Enqueue(3)
	value, ok := queue.PeekNext()
	if value != 5 {
		t.Error("Peeked value should be 5, not", value)
	}
	if !ok {
		t.Error("You should peek on empty queue", ok)
	}
	if queue.Len() != 2 {
		t.Error("Queue length should be 2, not", queue.Len())
	}
}
func TestPeekLastError(t *testing.T) {
	queue := NewQueue[int](3)
	_, ok := queue.PeekLast()
	if ok {
		t.Error("You can't peek on empty queue")
	}
}

func TestPeekLastValid(t *testing.T) {
	queue := NewQueue[int](3)
	queue.Enqueue(5)
	queue.Enqueue(3)
	value, ok := queue.PeekLast()
	if value != 3 {
		t.Error("Peeked value should be 3, not", value)
	}
	if !ok {
		t.Error("You should peek on empty queue", ok)
	}
	if queue.Len() != 2 {
		t.Error("Queue length should be 2, not", queue.Len())
	}
}
