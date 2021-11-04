package treelevels

func helper(root *Node, level int, result [][]string) [][]string {
	if root == nil {
		return result
	}

	for level >= len(result) {
		result = append(result, []string{})
	}

	result[level] = append(result[level], root.Val)

	if root.Left != nil {
		result = helper(root.Left, level+1, result)
	}

	if root.Right != nil {
		result = helper(root.Right, level+1, result)
	}

	return result
}

func DepthFirstRecursive(root *Node) [][]string {
	result := make([][]string, 0)

	return helper(root, 0, result)
}
