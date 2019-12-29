package node 

import (
	"testing"
)

func TestInit(t *testing.T) {
	initNode := New(23)
	if initNode.Item() != 23 {
		t.Errorf("expect 23 got %v \n", initNode.Item())
	}

	if initNode.Next() != nil {
		t.Errorf("expect nil got %v \n", initNode.Next())
	}
}

func TestRun(t *testing.T) {
	n := Node{23, nil}
	if n.Item() != 23 {
		t.Errorf("expect 23 got %v \n", n.Item())
	}

	if n.Next() != nil {
		t.Errorf("expect nil got %v \n", n.Next())
	}
}