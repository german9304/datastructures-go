/**
* Testing file for linked list
**/

package linkedListTesting

import (
	"testing"
	// "github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	ll "linkedList/main.go/linkedlist"
)

type appendSuite struct {
	suite.Suite
}

var alr ll.LinkedList 


func (s *appendSuite) BeforeTest(suiteName, testName string) {
	alr = ll.LinkedList{}

	alr.Append(2)
	alr.Append(3)
	alr.Append(4)
	alr.Append(10)

	// fmt.Printf("TestName: %v \n", testName)
}

func (s *appendSuite) AfterTest(suiteName, testName string) {

	// fmt.Println("This runs after test")
}


// Test for FindItem function
func (s *appendSuite) TestAppendToHead() {
	
	alr.AppendToHead(3)
	head := alr.Head()
	s.Equal(head.Item, 3, "head does not equal 3")
}

func  TestLinkedListAppendSuite(t *testing.T) {
	st := new(appendSuite)
	suite.Run(t, st)
}
