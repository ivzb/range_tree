package range_tree

const (
	Black = iota
	Red
)

type color int

type node struct {
	parent *node
	color  color

	left  *node
	right *node

	key   int64
	value interface{}
}

func (x *node) setParent(parent *node) {
	if x == nil {
		return
	}

	x.parent = parent
}

func (x *node) setColor(color color) {
	if x == nil {
		return
	}

	x.color = color
}

func (x *node) setLeft(left *node) {
	if x == nil {
		return
	}

	x.left = left
}

func (x *node) setRight(right *node) {
	if x == nil {
		return
	}

	x.right = right
}

func (x *node) isRed() bool {
	if x == nil {
		return false
	}

	return x.color == Red
}

func (x *node) isBlack() bool {
	if x == nil {
		return true
	}

	return x.color == Black
}

func (x *node) inorderTreeWalk(output []int64) []int64 {
	if x != nil {
		output = x.left.inorderTreeWalk(output)
		output = append(output, x.key)
		output = x.right.inorderTreeWalk(output)
	}

	return output
}

// O(h) time on a tree of height h
func (x *node) min() *node {
	if x == nil {
		return x
	}

	// reach the leftmost leaf
	for {
		if x.left == nil {
			break
		}

		x = x.left
	}

	return x
}

// O(h) time on a tree of height h
func (x *node) max() *node {
	if x == nil {
		return x
	}

	// reach the rightmost leaf
	for {
		if x.right == nil {
			break
		}

		x = x.right
	}

	return x
}

// O(h) time on a tree of height h
func (x *node) successor() *node {
	if x.right != nil {
		// the successor of x is just the leftmost node in x's right subtree
		return x.right.min()
	}

	// simply go up the tree from x until encounter
	// a node that is the left child of its parent
	y := x.parent

	for {
		if y == nil || x != y.right {
			break
		}

		x = y
		y = y.parent
	}

	return y
}
