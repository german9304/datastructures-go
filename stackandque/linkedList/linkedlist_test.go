package linkedList

import (
	// "log"
	"fmt"
	"testing"
	it "stackqueue/iterator"
)

func TestInsert(t *testing.T) {
	expectedRoot := 10
	li := New(nil, nil)
	li.Insert(34)
	li.Insert(44)
	li.Insert(10)

	item := li.Head().Item()
	if item != expectedRoot {
		t.Errorf("got %v , expected %v \n", item, expectedRoot)
	}
}

func TestAppend(t *testing.T) {
	li := New(nil, nil)
	li.Append(32)
	li.Append(50)
	li.Append(100)

	head := func() int {
		return li.Head().Item()
	}

	if head() != 32 {
		t.Errorf("got %v, expected %v  \n", head(), 32)
	}

	li.RemoveFront()

	if head() != 50 {
		t.Errorf("got %v, expected %v  \n", head(), 50)
	}

	fmt.Println("iterator append")
	var iter it.IIterator = li.Iterator()
	for iter.HasNext() {
		fmt.Printf("%v\n", iter.Next().Item())
	}
	fmt.Println("iterator append")
}

func TestRemoveFront(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6}
	li2 := New(nil, nil)

	if li2.RemoveFront() != nil {
		t.Errorf("should be nil")
	}


	li := New(nil, nil)
	for _, number := range numbers {
		li.Insert(number)
	}
	li.RemoveFront()

	// New Head 
	if got := li.Head().Item(); got != 5 {
		t.Errorf("got %v , expected %v \n", got, 5)
	}
	var iter it.IIterator = li.Iterator()
	for iter.HasNext() {
		fmt.Printf("%v\n", iter.Next().Item())
	}
}


func ExamplePrint() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	li := New(nil, nil)
	for _, number := range numbers {
		li.Insert(number)
	}
	li.Print()
	// output:
	// 6
	// 5
	// 4 
	// 3
	// 2
	// 1
}
