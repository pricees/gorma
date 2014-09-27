package btree

import (
	_ "fmt"
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

func equals(a, b IndexableList) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestSplice(t *testing.T) {

	vals := IndexableList{MyInt(1), MyInt(2), MyInt(3), MyInt(4)}

	var tests = []struct {
		exp IndexableList
		res func() IndexableList
	}{
		{
			// Remove nothing
			vals,
			func() IndexableList { return vals.Splice(0, 0) },
		},
		{
			// Remove 1
			IndexableList{MyInt(1), MyInt(2), MyInt(4)},
			func() IndexableList { return vals.Splice(2, 1) },
		},
		{
			// Remove 2
			IndexableList{MyInt(1), MyInt(4)},
			func() IndexableList { return vals.Splice(1, 2) },
		},
		{
			// Remove 2, Add 3
			IndexableList{MyInt(1), MyInt(101), MyInt(13), MyInt(14), MyInt(4)},
			func() IndexableList { return vals.Splice(1, 2, MyInt(101), MyInt(13), MyInt(14)) },
		},
	}

	for _, test := range tests {
		exp := test.exp
		res := test.res()
		if !equals(exp, res) {
			t.Errorf("expected %v, got %v", exp, res)
		}
	}
}
