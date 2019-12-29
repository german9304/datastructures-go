package main 


import (
	"fmt"
	"bst/main/node"
)

type bst struct {
	root *node.TreeNode
}

func (b *bst) insertNode(item int) {
	n := &node.TreeNode{item, nil, nil}
	  
	if b.root == nil {
		b.root = n
		return
	}
	b.insert(b.root, item)
}

func (b *bst) insert(root *node.TreeNode, item int) {

	if root.Item > item {
		if root.Left == nil {
			n := &node.TreeNode{item, nil, nil}
			fmt.Printf("inserting node %v \n", item)
			root.Left = n
		}else {
			b.insert(root.Left, item)
		}
	}

	if root.Item < item {
		if root.Right == nil {
			n := &node.TreeNode{item, nil, nil}
			fmt.Printf("inserting node %v \n", item)
			root.Right = n
		}else {
			b.insert(root.Right, item)
		}
	}
}

func (b *bst) removeItem(item int) {
	
}

func (b *bst) findItem(item int) {

}

func (b *bst) minNumber(item int) {
	
}

func (b *bst) printElements() {
	b.traverse(b.root)
}

func (b *bst) traverse(root *node.TreeNode) {
	if root != nil {
		b.traverse(root.Left)
		fmt.Printf("Element is: %v \n", root.Item)
		b.traverse(root.Right)
	}
}

func main() {
	bst := bst{nil}
	bst.insertNode(20)
	bst.insertNode(50)
	bst.insertNode(10)
	bst.insertNode(100)
	bst.printElements()
}