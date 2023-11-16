package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// If both left and right subtrees are non-nil
	if root.Left != nil && root.Left.Data > root.Data {
		return false
	}

	if root.Right != nil && root.Right.Data < root.Data {
		return false
	}

	// Recursively check both left and right subtrees
	if !(BTreeIsBinary(root.Left)) || !(BTreeIsBinary(root.Right)) {
		return false
	}

	return true
}
