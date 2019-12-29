package bst

import (
	"fmt"
	// "log"

	"bst/node"
)

type Bst struct {
	root    *node.TreeNode
	current *node.TreeNode
}

func (b *Bst) Root() *node.TreeNode {
	return b.root
}

func (b *Bst) Right() *node.TreeNode {
	return b.root.Right
}

func (b *Bst) Insert(item int) {
	newNode := &node.TreeNode{item, nil, nil}
	if b.root == nil {
		b.root = newNode
		b.current = newNode
	} else {
		b.current = b.root
	}
	b.insertNode(b.root, item)

}

// Inserts node into Binary Tree
func (b *Bst) insertNode(root *node.TreeNode, item int) {

	if root == nil {
		newNode := &node.TreeNode{item, nil, nil}
		if b.current.Item > item {
			b.current.Left = newNode
		} else {
			b.current.Right = newNode
		}

		return
	}

	if root.Item > item {
		b.current = root
		b.insertNode(root.Left, item)
	}

	if root.Item < item {
		b.current = root
		b.insertNode(root.Right, item)
	}
}

func (b *Bst) FindItem(item int) (*node.TreeNode, bool) {
	return nil, false
}


func (b *Bst) LeftMost(item int) (*node.TreeNode, bool) {
	root := b.Root()
	leftMostNode := b.innerLeftMost(root, item)
	if leftMostNode != nil {
		return leftMostNode, true
	}
	return leftMostNode, false
}



func (b *Bst) innerLeftMost(root *node.TreeNode, item int) *node.TreeNode {
	return nil
}

func (b *Bst) Remove(item int) {

}

func (b *Bst) Update(item int) {

}

func (b *Bst) InOrder() {

	root := b.root
	b.PrintInOder(root)
}

func (b *Bst) PrintInOder(node *node.TreeNode) {

	if node == nil {
		return
	}
	b.PrintInOder(node.Left)
	fmt.Println(node.Item)
	b.PrintInOder(node.Right)

}

func (b *Bst) PreOrder() {

}
