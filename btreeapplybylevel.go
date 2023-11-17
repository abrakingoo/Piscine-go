package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		base := queue[0]
		queue = queue[1:]

		f(base.Data)

		if base.Left != nil {
			queue = append(queue, base.Left)
		}
		if base.Right != nil {
			queue = append(queue, base.Right)
		}
	}
}
