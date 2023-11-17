package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	// If the node to be replaced is the root of the tree
	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		// If the node is the left child of its parent
		node.Parent.Left = rplc
	} else {
		// If the node is the right child of its parent
		node.Parent.Right = rplc
	}

	// If the replacement node is not nil, update its parent
	if rplc != nil {
		rplc.Parent = node.Parent
	}

	// Disconnect the replaced node from the tree
	node.Parent = nil
	node.Left = nil
	node.Right = nil

	return root
}
