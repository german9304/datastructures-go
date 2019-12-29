package bst 

import (
	"testing"
	"log"
)


func TestInsert(t *testing.T) {
	bst := Bst{nil, nil}
	bst.Insert(3)
	// bst.InOrder()
}


// Right most element
func TestLeftMost(t *testing.T) {
	bst := Bst{nil, nil}
	numbersToInsert := []int{10, 3, 14, 2, 5, 4}
	for _, v := range numbersToInsert {
		bst.Insert(v)
	}

	node, existsLeftMost := bst.LeftMost(10)

	if existsLeftMost == true {
		log.Printf("item %v \n", node.Item)
	}

}

// Left most element
func TestRightMost(t *testing.T) {

}


func TestRemove(t *testing.T) {
	bst := Bst{nil, nil}
	numbersToInsert := []int{10, 3, 14, 2, 5, 4}
	for _, v := range numbersToInsert {
		bst.Insert(v)
	}
	bst.InOrder()
}


// Checks wether print in order is correct
func ExampleInOrder() {
	bst := Bst{nil, nil}
	numbers := []int{10, 3, 4, 5, 19, 2, 7, 1, 34, 18}
	for _, v := range numbers {
		bst.Insert(v)
	}
	bst.InOrder()
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5 
	// 7
	// 10
	// 18
	// 19
	// 34
}