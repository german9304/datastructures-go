package queue

import (
	"stackqueue/linkedList"
	"stackqueue/node"
	it "stackqueue/iterator"
)


type IQueue interface {
	Enqueue(item int) *node.Node
	Dequeue() *node.Node
	Peek() *node.Node
	Iterator() it.IIterator
}

type Queue struct {
	li *linkedList.LinkedList
}


func New() *Queue {
	list := linkedList.New(nil, nil)
	return &Queue{list}
}

func (q *Queue) Iterator()  it.IIterator {
	return q.li.Iterator()
}

// Top of the list
func (q *Queue) Peek() *node.Node {
	return q.li.Head()
}
// Enqueues element
func (q *Queue) Enqueue(item int) *node.Node {
	newNode := q.li.Append(item)
	return newNode
}

// Dequeue element
func (q *Queue) Dequeue() *node.Node {
	removedNode := q.li.RemoveFront()
	return removedNode
}
