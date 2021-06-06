package cartesian

import "math"

type Tree struct {
	root *Node
}

func (tree *Tree) Root() *Node {
	return tree.root
}

type Node struct {
	value       interface{}
	idx         int
	from, to    int
	left, right *Node
	parent      *Node
}

func (n *Node) To() int {
	return n.to
}

func (n *Node) From() int {
	return n.from
}

func (n *Node) Idx() int {
	return n.idx
}

func (n *Node) Value() interface{} {
	return n.value
}
func (n *Node) Left() *Node {
	return n.left
}

func (n *Node) Right() *Node {
	return n.right
}

type Compare func(a interface{}, b interface{}) int

func NewTree(arr []interface{}, c Compare) (tree *Tree) {
	n := len(arr)
	if n == 0 {
		return &Tree{
			root: nil,
		}
	}

	cur := &Node{
		value: arr[0],
		idx:   0,
	}
	dummy := &Node{
		value:  -1,
		idx:    -1,
		right:  cur,
		parent: nil,
	}
	cur.parent = dummy
	for i := 1; i < n; i++ {
		e := arr[i]
		switch r := c(e, cur.value); {
		case r < 0:
			//	find parent <e or ==dummy
			for cur != dummy && c(e, cur.value) < 0 {
				cur = cur.parent
			}
			// insert right
			tmp := cur.right
			cur.right = &Node{
				value:  e,
				idx:    i,
				parent: cur,
				left:   tmp,
			}
			tmp.parent = cur.right
			cur = cur.right
		case r >= 0:
			//	lower
			cur.right = &Node{
				value:  e,
				idx:    i,
				parent: cur,
			}
			cur = cur.right
		}
	}
	traverseAndSetRange(dummy.right, 0, n-1)
	return &Tree{
		root: dummy.right,
	}
}

func NewIntTree(arr []int, minOnTop bool) (tree *Tree) {
	n := len(arr)
	if n == 0 {
		return &Tree{
			root: nil,
		}
	}
	dummy := &Node{
		idx:    -1,
		right:  nil,
		parent: nil,
	}
	if minOnTop {
		dummy.value = math.MinInt64
	} else {
		dummy.value = math.MaxInt64
	}
	cur := dummy
	for i := 0; i < n; i++ {
		e := arr[i]
		switch c := compareInt(e, cur.value.(int), minOnTop); {
		case c < 0:
			//	higher
			//find parent which is higher than e
			for compareInt(e, cur.value.(int), minOnTop) < 0 {
				cur = cur.parent
			}
			//	insert right
			tmp := cur.right
			cur.right = &Node{
				value:  e,
				idx:    i,
				parent: cur,
				left:   tmp,
			}
			tmp.parent = cur.right
			cur = cur.right
		case c >= 0:
			//	lower
			cur.right = &Node{
				value:  e,
				idx:    i,
				parent: cur,
			}
			cur = cur.right
		}
	}
	traverseAndSetRange(dummy.right, 0, n-1)

	return &Tree{
		root: dummy.right,
	}
}

func traverseAndSetRange(node *Node, from int, to int) {
	if node == nil {
		return
	}
	node.from = from
	node.to = to
	traverseAndSetRange(node.left, from, node.idx-1)
	traverseAndSetRange(node.right, node.idx+1, to)
}

func compareInt(a int, b int, minOnTop bool) int {
	if a == b {
		return 0
	}
	c := -1
	if a > b {
		c = 1
	}
	if !minOnTop {
		c = -c
	}
	return c
}
