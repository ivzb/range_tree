package range_tree

type Rbt struct {
	root *node
}

func NewRedBlackTree() *Rbt {
	return &Rbt{}
}

// O(n) time to walk an n-node RBT
func (rbt *Rbt) inorderTreeWalk() []int64 {
	return rbt.root.inorderTreeWalk([]int64{})
}

// O(lg n) time on a tree of height n
func (rbt *Rbt) Search(target int64) *node {
	child := rbt.root

	for {
		if child == nil || child.key == target {
			break
		}

		if child.key > target {
			child = child.left
		} else {
			child = child.right
		}
	}

	return child
}

// O(lg n)
func (rbt *Rbt) Insert(key int64, value interface{}) {
	child := &node{
		key:   key,
		value: value,
	}

	var parent *node
	temp := rbt.root

	for {
		if temp == nil {
			break
		}

		parent = temp

		if child.key < parent.key {
			temp = temp.left
		} else {
			temp = temp.right
		}
	}

	child.setParent(parent)

	if parent == nil {
		rbt.root = child
	} else if child.key < parent.key {
		parent.setLeft(child)
	} else {
		parent.setRight(child)
	}

	child.setLeft(nil)
	child.setRight(nil)
	child.setColor(Red)

	rbt.insertFixup(child)
}

func (rbt *Rbt) insertFixup(child *node) {
	for {
		if child.parent.isBlack() {
			break
		}

		if child.parent == child.parent.parent.left { // parent is left child of grand-parent
			uncle := child.parent.parent.right

			if uncle.isRed() {
				child.parent.setColor(Black)
				uncle.setColor(Black)
				child.parent.parent.setColor(Red)
				child = child.parent.parent
			} else {
				if child == child.parent.right {
					child = child.parent
					rbt.rotateLeft(child)
				}

				child.parent.setColor(Black)
				child.parent.parent.setColor(Red)
				rbt.rotateRight(child.parent.parent)
			}
		} else { // parent is right child of grand-parent
			uncle := child.parent.parent.left

			if uncle.isRed() {
				child.parent.setColor(Black)
				uncle.setColor(Black)
				child.parent.parent.setColor(Red)
				child = child.parent.parent
			} else {
				if child == child.parent.left {
					child = child.parent
					rbt.rotateRight(child)
				}

				child.parent.setColor(Black)
				child.parent.parent.setColor(Red)
				rbt.rotateLeft(child.parent.parent)
			}
		}
	}

	rbt.root.setColor(Black)
}

// Deleting a node z
//
// a) Node z has no left child. Replace z by its right child r, which may or may not be nil.
//
//     q         q
//     |         |
//     z    ->   r
//    / \       / \
//  nil  r
//      / \
//
// --------------------------------------------------------------------
// b) Node z has a left child l but no right child. Replace z by l
//
//     q          q
//     |          |
//     z     ->   l
//    / \        / \
//   l  nil
//  / \
//
// --------------------------------------------------------------------
// c) Node z has two children; its left child is node l, its right child is its successor y,
// and y's right child is node x. Replace z by y, updating y's left child to become l, but leaving x and y's right child.
//
//      q              q
//      |              |
//      z       ->     y
//    /   \           / \
//   l     y         l   x
//  / \   / \       / \ / \
//       nil x
//          / \
//
// --------------------------------------------------------------------
// d) Node z has two children (left child l and right child r), and its successor y != r
// lies withing the subtree rooted at r. Replace y by its own right child x, and set y to be r's parent.
// Then, set y to be q's child and the parent of l.
//
//       q            q                        q
//       |            |                        |
//       z      ->    z       y        ->      y
//     /   \         /       / \             /   \
//    l     r       l      nil  r           l     r
//   / \   / \     /           / \         / \   / \
//        y                   x                 x
//       / \                 / \               / \
//     nil  x
//         / \
//
// O(lg n)
func (rbt *Rbt) delete(z *node) {
	y := z
	yOriginalColor := y.color

	var x *node

	if z.left == nil {
		x = z.right
		rbt.transplant(z, z.right)
	} else if z.right == nil {
		x = z.left
		rbt.transplant(z, z.left)
	} else {
		y = z.right.min()
		yOriginalColor = y.color
		x = y.right

		if y.parent == z {
			x.setParent(y)
		} else {
			rbt.transplant(y, y.right)
			y.setRight(z.right)
			y.right.setParent(y)
		}

		rbt.transplant(z, y)
		y.setLeft(z.left)
		y.left.setParent(y)
		y.setColor(z.color)
	}

	if yOriginalColor == Black {
		rbt.deleteFixup(x)
	}
}

// O(lg n)
func (rbt *Rbt) deleteFixup(x *node) {
	// while(x != rbt.root && x.color == black)
	for {
		if x == rbt.root || x.isRed() {
			break
		}

		if x == x.parent.left {
			w := x.parent.right // todo: rename to sibling

			// x's sibling is red. Switch the colors of w and
			// x.parent and then perform a left-rotation on x.parent
			if w.isRed() {
				w.setColor(Black)
				x.parent.setColor(Red)
				rbt.rotateLeft(x.parent)
				w = x.parent.right
			}

			// x's sibling w is black; both of w's children
			// are black. take one black off x and w; repeat the while
			// loop with p[x], adding an extra black to p[x]. if we
			// entered through previous case, x is red and the loop terminates.
			if w.left.isBlack() && w.right.isBlack() {
				w.setColor(Red)
				x = x.parent
			} else {
				// w is black, left[w] is red, right[w] is black.
				// Switch the colors of w and left[w] and right-rotate w
				if w.right.isBlack() {
					w.left.setColor(Black)
					w.setColor(Red)
					rbt.rotateRight(w)
					w = x.parent.right
				}

				w.setColor(x.parent.color)
				x.parent.setColor(Black)
				w.right.setColor(Black)
				rbt.rotateLeft(x.parent)
				x = rbt.root
			}
		} else {
			w := x.parent.left

			if w.isRed() {
				w.setColor(Black)
				x.parent.setColor(Red)
				rbt.rotateRight(x.parent)
				w = x.parent.left
			}

			if w.right.isBlack() && w.left.isBlack() {
				w.setColor(Red)
				x = x.parent
			} else {
				if w.left.isBlack() {
					w.right.setColor(Black)
					w.setColor(Red)
					rbt.rotateLeft(w)
					w = x.parent.left
				}

				w.setColor(x.parent.color)
				x.parent.setColor(Black)
				w.left.setColor(Black)
				rbt.rotateRight(x.parent)
				x = rbt.root
			}
		}
	}

	x.setColor(Black)
}

func (rbt *Rbt) transplant(u, v *node) {
	// replaces one subtree as a child of its parent with another subtree
	// when subtree gets replaced at node u with the subree rooted at node v,
	// node u's parent becomes node v's parent
	if u.parent == nil {
		rbt.root = v
	} else if u == u.parent.left { // u is a left child
		u.parent.left = v
	} else { // u is a right child
		u.parent.right = v
	}

	v.parent = u.parent
}

// Rotations
//
//      |                              |
//      y        left-rotation         x
//     / \       <-------------       / \
//    x   c                          a   y
//   / \         right-rotation         / \
//  a   b        ------------->        b   c
//
// O(1)
func (rbt *Rbt) rotateLeft(x *node) {
	y := x.right
	x.right = y.left

	if y.left != nil {
		y.left.parent = x
	}

	y.parent = x.parent

	if x.parent == nil {
		rbt.root = y
	} else if x == x.parent.left {
		x.parent.left = y
	} else {
		x.parent.right = y
	}

	y.left = x
	x.parent = y
}

// O(1)
func (rbt *Rbt) rotateRight(y *node) {
	x := y.left
	y.left = x.right

	if x.right != nil {
		x.right.parent = y
	}

	x.parent = y.parent

	if y.parent == nil {
		rbt.root = x
	} else if y == y.parent.right {
		y.parent.right = x
	} else {
		y.parent.left = x
	}

	x.right = y
	y.parent = x
}
