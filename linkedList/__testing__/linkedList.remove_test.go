/**
* Testing file for linked list
**/

package linkedListTesting

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	ll "linkedList/main.go/linkedlist"
	"fmt"
)

type removeSuite struct {
	suite.Suite
}

var al ll.LinkedList 

func (s *removeSuite) BeforeTest(suiteName, testName string) {
	al = ll.LinkedList{}

	al.Append(2)
	al.Append(3)
	al.Append(4)
	al.Append(10)

	// fmt.Printf("TestName: %v \n", testName)
}

func (s *removeSuite) AfterTest(suiteName, testName string) {

	// fmt.Println("This runs after test")
}

// Test for FindItem function
func (s *removeSuite) TestFindItem() {
	_, f1 := al.FindItem(2)
	_, f2 := al.FindItem(3)
	_, f3 := al.FindItem(4)
	_, f4 := al.FindItem(10)
	item5, f5 := al.FindItem(11)
	// Checks for items
	s.Equal(f1, true, "Item1 should be 2")
	s.Equal(f2, true, "Item2 should be 3")
	s.Equal(f3, true, "Item3 should be 4")
	s.Equal(f4, true, "Item4 should be 10")
	// Checks if item was not found
	s.Equal(f5, false, "f5 should be false")
	// is item5 nil? means it was not found
	if s.Nil(item5) {
		fmt.Println("Item5 is nil because it was not found")
	}
}

// Function test when head is removed
func (s *removeSuite) TestRemoveHead() {
	al.RemoveItem(2)
	prevHead, exists := al.FindItem(2)
	s.Equal(al.Head().Item, 3, "Head does not equal 3")
	s.Equal(exists, false, "prev head does not exists , return false")
	if s.Nil(prevHead) {
		fmt.Println("previous Head is nil, it does not longer exists")
	}
}
// Function test when tail is removed 
func (s *removeSuite) TestRemoveTail() {
	removed := al.RemoveItem(10)
	prevTail, found := al.FindItem(10)
	s.Equal(removed, true, "Element was not removed")
	s.Equal(found, false, "Element was not removed")
	if s.Nil(prevTail) {
		fmt.Println("previous tail is nil, it does not longer exists")
	}
}
// Function test when middle node is removed
func (s *removeSuite) TestRemoveMiddle() {
	prevNode,_ := al.FindItem(3)
	afterNode,_ := al.FindItem(10)
	removedNode := al.RemoveItem(4)
	findPrevHead, found := al.FindItem(4)
	s.Equal(removedNode, true, "Node should be removed")
	s.Equal(prevNode.Next, afterNode, "Prev node does not equal after node")
	s.Equal(found, false, "Prev node does not equal after node")
	if s.Nil(findPrevHead) {
		fmt.Println("previous node is nil, it does not longer exists")
	}
	
}

func  TestLinkedListSuite(t *testing.T) {
	st := new(removeSuite)
	suite.Run(t, st)
}
