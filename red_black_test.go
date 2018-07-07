package range_tree

import (
	"testing"

	"github.com/ivzb/tick/tick/assert"
)

func TestRedBlackTree_inorderTreeWalk(t *testing.T) {
	rbt := &Rbt{}
	rbt.root = &node{
		key: 5,
		left: &node{
			key: 4,
		},
		right: &node{
			key: 6,
		},
	}

	expected := []int64{4, 5, 6}
	actual := rbt.inorderTreeWalk()

	assert.Equal(t, expected, actual)
}

func TestRedBlackTree_insert(t *testing.T) {
	root := &node{key: 11, color: Black}

	root.left = &node{parent: root, key: 2, color: Red}
	root.left.left = &node{parent: root.left, key: 1, color: Black}
	root.left.right = &node{parent: root.left, key: 7, color: Black}
	root.left.right.left = &node{parent: root.left.right, key: 5, color: Red}
	root.left.right.right = &node{parent: root.left.right, key: 8, color: Red}

	root.right = &node{parent: root, key: 14, color: Black}
	root.right.right = &node{parent: root.right, key: 15, color: Red}

	rbt := &Rbt{}
	rbt.root = root

	assert.Equal(t, rbt.root.key, int64(11))
	assert.Equal(t, rbt.root.color, color(Black))

	assert.Equal(t, rbt.root.left.key, int64(2))
	assert.Equal(t, rbt.root.left.color, color(Red))
	assert.Equal(t, rbt.root.left.left.key, int64(1))
	assert.Equal(t, rbt.root.left.left.color, color(Black))
	assert.Equal(t, rbt.root.left.right.key, int64(7))
	assert.Equal(t, rbt.root.left.right.color, color(Black))
	assert.Equal(t, rbt.root.left.right.left.key, int64(5))
	assert.Equal(t, rbt.root.left.right.left.color, color(Red))
	assert.Equal(t, rbt.root.left.right.right.key, int64(8))
	assert.Equal(t, rbt.root.left.right.right.color, color(Red))

	assert.Equal(t, rbt.root.right.key, int64(14))
	assert.Equal(t, rbt.root.right.color, color(Black))
	assert.Equal(t, rbt.root.right.right.key, int64(15))
	assert.Equal(t, rbt.root.right.right.color, color(Red))

	rbt.Insert(4, 4)

	assert.Equal(t, rbt.root.key, int64(7))
	assert.Equal(t, rbt.root.color, color(Black))

	assert.Equal(t, rbt.root.left.key, int64(2))
	assert.Equal(t, rbt.root.left.color, color(Red))
	assert.Equal(t, rbt.root.left.left.key, int64(1))
	assert.Equal(t, rbt.root.left.left.color, color(Black))
	assert.Equal(t, rbt.root.left.right.key, int64(5))
	assert.Equal(t, rbt.root.left.right.color, color(Black))
	assert.Equal(t, rbt.root.left.right.left.key, int64(4))
	assert.Equal(t, rbt.root.left.right.left.color, color(Red))

	assert.Equal(t, rbt.root.right.key, int64(11))
	assert.Equal(t, rbt.root.right.color, color(Red))
	assert.Equal(t, rbt.root.right.left.key, int64(8))
	assert.Equal(t, rbt.root.right.left.color, color(Black))
	assert.Equal(t, rbt.root.right.right.key, int64(14))
	assert.Equal(t, rbt.root.right.right.color, color(Black))
	assert.Equal(t, rbt.root.right.right.right.key, int64(15))
	assert.Equal(t, rbt.root.right.right.right.color, color(Red))
}

// Remove 7b
//
//      5b                  5b
//    /   \               /   \
//   3b    7b    -->     3b    8b
//  / \   / \           / \   /
// 2r  4r6r  8r        2r  4r6r
//
//
//
func TestRedBlackTree_Delete_black_sibling(t *testing.T) {
	rbt := &Rbt{}
	rbt.Insert(5, 5)

	rbt.Insert(3, 3)
	rbt.Insert(7, 7)

	rbt.Insert(2, 2)
	rbt.Insert(6, 6)

	rbt.Insert(4, 4)
	rbt.Insert(8, 8)

	// root
	assert.Equal(t, rbt.root.key, int64(5))
	assert.Equal(t, rbt.root.color, color(Black))

	// left
	assert.Equal(t, rbt.root.left.key, int64(3))
	assert.Equal(t, rbt.root.left.color, color(Black))
	assert.Equal(t, rbt.root.left.left.key, int64(2))
	assert.Equal(t, rbt.root.left.left.color, color(Red))
	assert.Equal(t, rbt.root.left.right.key, int64(4))
	assert.Equal(t, rbt.root.left.right.color, color(Red))

	// right
	assert.Equal(t, rbt.root.right.key, int64(7))
	assert.Equal(t, rbt.root.right.color, color(Black))
	assert.Equal(t, rbt.root.right.left.key, int64(6))
	assert.Equal(t, rbt.root.right.left.color, color(Red))
	assert.Equal(t, rbt.root.right.right.key, int64(8))
	assert.Equal(t, rbt.root.right.right.color, color(Red))

	rbt.delete(rbt.root.right)

	// root
	assert.Equal(t, rbt.root.key, int64(5))
	assert.Equal(t, rbt.root.color, color(Black))

	// left
	assert.Equal(t, rbt.root.left.key, int64(3))
	assert.Equal(t, rbt.root.left.color, color(Black))
	assert.Equal(t, rbt.root.left.left.key, int64(2))
	assert.Equal(t, rbt.root.left.left.color, color(Red))
	assert.Equal(t, rbt.root.left.right.key, int64(4))
	assert.Equal(t, rbt.root.left.right.color, color(Red))

	// right
	assert.Equal(t, rbt.root.right.key, int64(8))
	assert.Equal(t, rbt.root.right.color, color(Black))
	assert.Equal(t, rbt.root.right.left.key, int64(6))
	assert.Equal(t, rbt.root.right.left.color, color(Red))
}

func TestRedBlackTree_empty(t *testing.T) {
	rbt := &Rbt{}

	var expected *node
	actual := rbt.Search(5)

	assert.Equal(t, expected, actual)
}

func TestRedBlackTree_single(t *testing.T) {
	rbt := &Rbt{}
	rbt.root = &node{key: 5, value: 5}

	expected := &node{key: 5, value: 5}
	actual := rbt.Search(5)

	assert.Equal(t, expected.value, actual.value)
}

func TestRedBlackTree_min_empty(t *testing.T) {
	var root *node
	actual := root.min() == nil

	assert.True(t, actual, "min empty")
}

func TestRedBlackTree_rotateLeft(t *testing.T) {
	root := &node{key: 7}
	root.right = &node{parent: root, key: 18}

	root.right.left = &node{parent: root.right, key: 11}
	root.right.left.left = &node{parent: root.right.left, key: 9}
	root.right.left.right = &node{parent: root.right.left, key: 14}
	root.right.left.right.left = &node{parent: root.right.left.right, key: 12}
	root.right.left.right.right = &node{parent: root.right.left.right, key: 17}

	root.right.right = &node{parent: root.right, key: 19}
	root.right.right.right = &node{parent: root.right.right, key: 22}
	root.right.right.right.left = &node{parent: root.right.right.right, key: 20}

	rbt := &Rbt{}
	rbt.root = root

	assert.Equal(t, root.right.key, int64(18))
	rbt.rotateLeft(rbt.root)
	assert.Equal(t, root.right.key, int64(11))
}

func TestRedBlackTree_rotateRight(t *testing.T) {
	root := &node{key: 7}
	root.right = &node{parent: root, key: 11}

	root.right.left = &node{parent: root.right, key: 18}
	root.right.left.left = &node{parent: root.right.left, key: 9}
	root.right.left.right = &node{parent: root.right.left, key: 14}
	root.right.left.right.left = &node{parent: root.right.left.right, key: 12}
	root.right.left.right.right = &node{parent: root.right.left.right, key: 17}

	root.right.right = &node{parent: root.right, key: 19}
	root.right.right.right = &node{parent: root.right.right, key: 22}
	root.right.right.right.left = &node{parent: root.right.right.left, key: 20}

	rbt := &Rbt{}
	rbt.root = root

	assert.Equal(t, root.right.key, int64(11))
	rbt.rotateRight(rbt.root.right)
	assert.Equal(t, root.right.key, int64(18))
}

func TestRedBlackTree_node_min(t *testing.T) {
	root := &node{
		key: 5,
		left: &node{
			key: 4,
		},
		right: &node{
			key: 6,
		},
	}

	expected := &node{key: 4}
	actual := root.min()

	assert.Equal(t, expected.key, actual.key)
}

func TestRedBlackTree_node_max(t *testing.T) {
	root := &node{
		key: 5,
		left: &node{
			key: 4,
		},
		right: &node{
			key: 6,
		},
	}

	expected := &node{key: 6}
	actual := root.max()

	assert.Equal(t, expected.key, actual.key)
}

func TestRedBlackTree_node_max_empty(t *testing.T) {
	var root *node
	actual := root.max() == nil

	assert.True(t, actual, "min empty")
}

func TestRedBlackTree_node_successfor_left(t *testing.T) {
	root := &node{key: 5}

	root.left = &node{parent: root, key: 3}
	root.left.left = &node{parent: root.left, key: 2}
	root.left.right = &node{parent: root.left, key: 4}

	root.right = &node{parent: root, key: 7}

	expected := &node{key: 3}
	actual := root.left.left.successor()

	assert.Equal(t, expected.key, actual.key)
}

func TestRedBlackTree_node_successfor_root(t *testing.T) {
	root := &node{key: 5}
	root.left = &node{parent: root, key: 4}
	root.right = &node{parent: root, key: 6}

	expected := &node{key: 6}
	actual := root.successor()

	assert.Equal(t, expected.key, actual.key)
}

func TestRedBlackTree_node_successfor_right(t *testing.T) {
	root := &node{key: 5}

	root.left = &node{parent: root, key: 4}

	root.right = &node{parent: root, key: 7}
	root.right.left = &node{parent: root.right, key: 6}
	root.right.right = &node{parent: root.right, key: 8}

	expected := &node{key: 7}
	actual := root.right.left.successor()

	assert.Equal(t, expected.key, actual.key)
}
