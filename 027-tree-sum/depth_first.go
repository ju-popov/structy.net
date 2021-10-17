package treesum

func DepthFirst(root *Node) int64 {
	if root == nil {
		return 0
	}

	return root.Value + DepthFirst(root.Left) + DepthFirst(root.Right)
}
