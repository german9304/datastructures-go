package queue

import (
	"testing"
	"fmt"
)

// Queue test
func TestQueue(t *testing.T) {
	var queue IQueue = New()
	queue.Enqueue(32)
	queue.Enqueue(40)
	queue.Enqueue(100)

	peek := queue.Peek().Item()
	if peek != 32 {
		t.Errorf("got %v expected \n", peek)
	}

	iter := queue.Iterator()
	for iter.HasNext() {
		fmt.Printf("item %v \n", iter.Next().Item())
	}
}

func TestDequeue(t *testing.T) {
	var queue IQueue = New()
	queue.Enqueue(32)
	queue.Enqueue(40)
	queue.Enqueue(100)
	queue.Enqueue(2)

	queue.Dequeue()

	fmt.Println("iterator Dequeue")
	iter := queue.Iterator()
	for iter.HasNext() {
		fmt.Printf("item %v \n", iter.Next().Item())
	}
	fmt.Println("iterator Dequeue")

	peek := queue.Peek()
	if peek.Item() != 40 {
		t.Errorf("wrong dequeue")
	}
	queue.Dequeue()

	peek = queue.Peek()
	if peek.Item() != 100 {
		t.Errorf("wrong dequeue")
	}
}