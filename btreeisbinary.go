package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// If both left and right subtrees are non-nil
	if root.Left != nil && root.Right != nil {
		return false
	}

	// Recursively check both left and right subtrees
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
