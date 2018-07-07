package range_tree

import (
	"fmt"
)

type Rt struct {
	Rbt
}

func NewRangeTree() *Rt {
	return &Rt{}
}

// O(lg n + k) to output all k nodes
func (rt *Rt) Search(low int64, high int64) ([]interface{}, error) {
	if low >= high {
		return nil, fmt.Errorf("low(%d) should be less than high(%d)", low, high)
	}

	// At some node in the tree, the search paths to low and high will diverge.
	// Let lca be the last common ancestor that these two search paths have.
	lca := rt.root

	for {
		if lca == nil {
			break
		}

		if low < lca.key && high < lca.key {
			lca = lca.left
			continue
		} else if low > lca.key && high > lca.key {
			lca = lca.right
			continue
		}

		break
	}

	if lca == nil {
		return []interface{}{}, nil
	}

	nodes := []*node{}

	if inRange(lca.key, low, high) {
		nodes = append(nodes, lca)
	}

	getLeftNodes(low, high, lca.left, &nodes)
	getRightNodes(low, high, lca.right, &nodes)

	values := []interface{}{}

	for _, node := range nodes {
		values = append(values, node.value)
	}

	return values, nil
}

func getLeftNodes(low int64, high int64, node *node, nodes *[]*node) {
	if node == nil {
		return
	}

	if inRange(node.key, low, high) {
		*nodes = append(*nodes, node)
	}

	if node.key >= low {
		// For every node in the search path from lca to low, if the key at
		// node is greater than low, append to nodes
		getAllNodes(node.right, nodes)
		getLeftNodes(low, high, node.left, nodes)
	} else {
		getLeftNodes(low, high, node.right, nodes)
	}
}

func getRightNodes(low int64, high int64, node *node, nodes *[]*node) {
	if node == nil {
		return
	}

	if inRange(node.key, low, high) {
		*nodes = append(*nodes, node)
	}

	if node.key <= high {
		// For every node in the search path from lca to high, if the key at
		// node is less than high, append to nodes
		getAllNodes(node.left, nodes)
		getRightNodes(low, high, node.right, nodes)
	} else {
		getRightNodes(low, high, node.left, nodes)
	}
}

func getAllNodes(root *node, nodes *[]*node) {
	if root == nil {
		return
	}

	*nodes = append(*nodes, root)
	getAllNodes(root.left, nodes)
	getAllNodes(root.right, nodes)
}

func inRange(key int64, low int64, high int64) bool {
	return low <= key && key <= high
}
