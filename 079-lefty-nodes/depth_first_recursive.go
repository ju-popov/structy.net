package leftynodes

func helper(root *Node, level int, result []string) []string {
	if root == nil {
		return result
	}

	if level == len(result) {
		result = append(result, root.Val)
	}

	if root.Left != nil {
		result = helper(root.Left, level+1, result)
	}

	if root.Right != nil {
		result = helper(root.Right, level+1, result)
	}

	return result
}

func DepthFirstRecursive(root *Node) []string {
	return helper(root, 0, make([]string, 0))
}
