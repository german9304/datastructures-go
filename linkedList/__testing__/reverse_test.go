package linkedListTesting

import (
	"testing"
	ll "linkedList/main.go/linkedlist"
// 	"fmt"
)


func TestReverse(t *testing.T) {
	alr = ll.LinkedList{}
	alr.Append(2)
	alr.Append(3)
	alr.Append(4)
	alr.Append(10)
	alr.ReverseNode(alr.Head())
	alr.PrintItems()
}