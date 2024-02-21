package piscine

func BTreeIsBinary(root *TreeNode) bool {
	var prev *TreeNode
	return isBST(root, &prev)
}

func isBST(node *TreeNode, prev **TreeNode) bool {
	if node == nil {
		return true
	}

	if !isBST(node.Left, prev) {
		return false
	}

	if *prev != nil && (*prev).Data >= node.Data {
		return false
	}

	*prev = node

	return isBST(node.Right, prev)
}
