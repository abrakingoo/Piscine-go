package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rcount := BTreeLevelCount(root.Right)
	lcont := BTreeLevelCount(root.Left)

	if rcount > lcont {
		return rcount + 1
	} else {
		return lcont + 1
	}
}
