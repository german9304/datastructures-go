package linkedlist

import (
	"linkedList/main.go/node"
	"fmt"
)

type LinkedList struct {
	head *node.Node
	curr *node.Node
	size int
}

// initialize linked list 
func (al *LinkedList) Init() {
	al.head = nil
	al.curr = nil
	al.size = 0
}
/**
*  Appends to head of linked list
* 
*/
func (al *LinkedList) AppendToHead(item int) {
	an := &node.Node{item, nil}

	if al.head == nil {
		al.head = an
		al.curr = an
	} else {
		currHead := al.head 
		al.head = an
		an.Next = currHead
	}
}

// Returns last node of the list 

func (al *LinkedList) LastNode() *node.Node{
	cur := al.head
	for cur != nil {
		if cur == nil {
			return cur
		}
		cur = cur.Next
	}
	return nil
}

// Appends element to linked list
func (al *LinkedList) Append(item int) {
	n := &node.Node{item, nil}
	if al.head == nil {
		al.head = n
		al.curr = n 
		al.size = al.size + 1
		return 
	}
	curr := al.curr
	curr.Next = n
	al.curr = n
	al.size = al.size + 1
}

// Finds item in the linked list, based on the parameter
func (al *LinkedList) FindItem(item int) (*node.Node, bool) {
	cur := al.head

	for cur != nil {
		if cur.Item == item {
			return cur, true
		}
		cur = cur.Next
	}
	return nil, false
}

// Prints items in the list
func (al *LinkedList) PrintItems() {
	head := al.head
	
	for head != nil {
		fmt.Println(head.Item)
		head = head.Next
	}
}

// Removes item in the list
func (al *LinkedList) RemoveItem(item int) bool{
	cur := al.head
	nextNode := cur.Next

	n, f := al.FindItem(item)
	if f == true && n == cur {
		al.head = nextNode
		al.size = al.size - 1
		return true
	} else {

		for nextNode != nil {
			if nextNode.Item == item {
				nNode := nextNode.Next
				cur.Next = nNode
				al.size = al.size - 1
				return true
			}
			cur = nextNode
			nextNode = nextNode.Next
		}
    }
	return false
}


func (al *LinkedList) Head() *node.Node {
	return al.head
}



func (al *LinkedList) ReverseNode(head *node.Node) *node.Node{

	if head.Next == nil {
		return head
	}
	 cur := al.ReverseNode(head.Next)

	if head.Next != nil {
		curNode := head
		nNode := head.Next
		nNode.Next = curNode
		nNode.Next.Next = nil
		if head == al.head {
			al.head = cur
		} 
	} 
	return cur
}

// 1 -> 2 -> 3 -> 4 -> nil