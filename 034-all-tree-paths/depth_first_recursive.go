package alltreepaths

func helper(root *Node, path []string) [][]string {
	path = append(path, root.Value)

	if (root.Left == nil) && (root.Right == nil) {
		// Important: Copy path array
		pathCopy := make([]string, len(path))
		copy(pathCopy, path)

		return [][]string{pathCopy}
	}

	var result [][]string

	if root.Left != nil {
		result = append(result, helper(root.Left, path)...)
	}

	if root.Right != nil {
		result = append(result, helper(root.Right, path)...)
	}

	return result
}

func DepthFirstRecursive(root *Node) [][]string {
	if root == nil {
		return [][]string{}
	}

	return helper(root, []string{})
}
