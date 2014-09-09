package btree

import (
	"testing"
)

func TestNodeIsLeaf(t *testing.T) {
	var n Node
	if n.isLeaf() != true {
		t.Error("New node should be leaf")
	}
}

func TestNodeIsFull(t *testing.T) {
	n := &Node{degree: 1}
	n.keys = append(n.keys, 1)

	if n.isFull() != true {
		t.Error("New node should be full")
	}
}

func TestSearchRoot(t *testing.T) {
}

func TestSplice(t *testing.T) {
}
