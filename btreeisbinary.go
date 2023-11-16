package piscine

func BTreeIsBinary(root *TreeNode) bool {
	if root == nil {
		return true
	}

	// If either left or right subtree is nil (not both)
	if (root.Left == nil && root.Right != nil) || (root.Left != nil && root.Right == nil) {
		return false
	}

	// Recursively check both left and right subtrees
	return BTreeIsBinary(root.Left) && BTreeIsBinary(root.Right)
}
