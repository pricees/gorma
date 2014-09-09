package btree

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
		b.root.split(0)
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

func (n Node) split(key int) []int {
	return []int{}
}

func (n Node) insert(key int, value string) {
	var i int
	for thisKey := range n.keys {
		if key <= thisKey {
			break
		}
		i += 1
	}

	if n.isLeaf() {
		splice(n.keys, i, 0, key)
		splice(n.values, i, 0, key)
	} else {
		child := n.children[i]

		if child.isFull() {
			n.split(i)
		}

		child.insert(key, value)
	}
}

type Interface interface{}

// Splice changes content of an array
// :array is array to change
// :startIndex, where to begin changes
// :howMany, old array elements to remove
// :... elements to add
func splice(array []Interface, startIndex int, howMany int, elements Interface) []Interface {
	startOfArray := array[:startIndex]
	startOfArray = append(startOfArray, elements)
	endOfArray := array[startIndex+howMany:]
	return append(startOfArray, endOfArray)
}

func (n Node) isLeaf() bool {
	return len(n.children) == 0
}

func (n Node) isFull() bool {
	return len(n.keys) >= 2*n.degree-1
}
