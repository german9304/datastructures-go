package iterator

import (
	"stackqueue/node"
)

type IIterator interface {
	HasNext() bool
	Next() *node.Node
}