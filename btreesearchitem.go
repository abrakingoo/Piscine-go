package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	if elem == root.Data {
		return root
	}
	if elem < root.Data {
		BTreeSearchItem(root.Left, elem)
		if elem == root.Left.Data {
			return root.Left
		}
	} else if elem > root.Data {
		BTreeSearchItem(root.Right, elem)
		if elem == root.Right.Data {
			return root.Right
		}
	}
	return nil
}
