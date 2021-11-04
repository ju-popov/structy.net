package treesum

func DepthFirstRecursive(root *Node) int64 {
	if root == nil {
		return 0
	}

	return root.Val + DepthFirstRecursive(root.Left) + DepthFirstRecursive(root.Right)
}
