package alltreepaths

func DepthFirstRecursive(root *Node) [][]string {
	results := make([][]string, 0)

	if root == nil {
		return results
	}

	if (root.Left == nil) && (root.Right == nil) {
		results = append(results, []string{root.Val})
		return results
	}

	if root.Left != nil {
		for _, value := range DepthFirstRecursive(root.Left) {
			results = append(results, append([]string{root.Val}, value...))
		}
	}

	if root.Right != nil {
		for _, value := range DepthFirstRecursive(root.Right) {
			results = append(results, append([]string{root.Val}, value...))
		}
	}

	return results
}
