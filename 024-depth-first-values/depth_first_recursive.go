package depthfirstvalues

func DepthFirstRecursive(root *Node) []string {
	result := []string{}

	if root == nil {
		return result
	}

	result = append(result, root.Val)
	result = append(result, DepthFirstRecursive(root.Left)...)
	result = append(result, DepthFirstRecursive(root.Right)...)

	return result
}
