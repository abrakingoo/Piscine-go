package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	new_node := &TreeNode{Data: data}

	if root == nil {
		return new_node
	}

	if new_node.Data > root.Data {
		root.Right = BTreeInsertData(new_node.Right, data)
		root.Right.Parent = root

	} else {
		root.Left = BTreeInsertData(root.Left, data)
		root.Left.Parent = root
	}

	return root
}
