package btree

import _ "fmt"

var (
	degree = 2
)

type BTree struct {
	degree int
	root   Node
}

func (b BTree) search(key int) []int {
	return b.root.search(key)
}

func (b BTree) insert(key int, value string) {
	if b.root.isFull() {
		oldRoot := b.root
		b.root = Node{degree: b.degree}
		b.root.children = append(b.root.children, oldRoot)
		//b.root.split(0)
	}

	b.root.insert(key, value)
}

type Node struct {
	degree   int
	children []Node
	keys     []int
	values   []string
}

func (n Node) search(key int) []int {
	return []int{}
}

/*
func (n Node) split(i int) {
	splitChild := n.children[i]
	newChild := new(Node)
	middle := degree - 1

	newChild.keys = splitChild.keys.splice(middle + 1)
	newChild.values = splitChild.values.splice(middle + 1)
	newChild.children = splitChild.children.splice(middle + 1)

	n.children.splice(i+1, 0, newChild)
	n.keys.splice(i, 0, splitChild.keys[middle])
	n.values.splice(i, 0, splitChild.values[middle])

	splitChild.keys.splice(middle)
	splitChild.values.splice(middle)
}

*/

func (n Node) insert(key int, value string) {
	var i int
	for thisKey := range n.keys {
		if key <= thisKey {
			break
		}
		i += 1
	}

	if n.isLeaf() {
		//splice(n.keys, i, 0, key)
		//splice(n.values, i, 0, key)
	} else {
		child := n.children[i]

		if child.isFull() {
			//			n.split(i)
		}

		child.insert(key, value)
	}
}

type Indexable interface {
	value()
}

type MyInt int

func (m MyInt) value() {
	//return int(m)
}

type MyString string

func (m MyString) value() {
	//return int(m)
}

type IndexableMap map[Indexable]string
type IndexableList []Indexable

// Insert inserts the value into the slice at the specified index,
// which must be in range.
// The slice must have room for the new element.
func Insert(slice []int, index, value int) []int {
	// Grow the slice by one element.
	slice = slice[0 : len(slice)+1]
	// Use copy to move the upper part of the slice out of the way and open a hole.
	copy(slice[index+1:], slice[index:])
	// Store the new value.
	slice[index] = value
	// Return the result.
	return slice
}

// Splice changes content of an array
// :startIndex, where to begin changes
// :howMany, old array elements to remove
// :... elements to add
func (il IndexableList) Splice(startIndex int, howMany int, elements ...Indexable) IndexableList {

	if howMany == 0 {
		return il
	}

	length := len(il) + len(elements) - howMany
	slice := make(IndexableList, length)

	copy(slice[:startIndex], il[:startIndex])

	var numNewElements int
	for i, element := range elements {
		if element != nil {
			numNewElements++
			slice[startIndex+i] = element
		}
	}
	copy(slice[startIndex+numNewElements:], il[(startIndex+howMany):])

	return slice
}

func (n Node) isLeaf() bool {
	return len(n.children) == 0
}

func (n Node) isFull() bool {
	return len(n.keys) >= 2*n.degree-1
}
