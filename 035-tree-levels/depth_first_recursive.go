package treelevels

func helper(root *Node, level int, result *[][]string) {
	if root == nil {
		return
	}

	for len(*result) <= level {
		*result = append(*result, []string{})
	}

	(*result)[level] = append((*result)[level], root.Value)

	if root.Left != nil {
		helper(root.Left, level+1, result)
	}

	if root.Right != nil {
		helper(root.Right, level+1, result)
	}
}

func DepthFirstRecursive(root *Node) [][]string {
	result := [][]string{}

	helper(root, 0, &result)

	return result
}
