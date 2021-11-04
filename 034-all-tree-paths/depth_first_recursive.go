package alltreepaths

func DepthFirstRecursive(root *Node) [][]string {
	result := make([][]string, 0)

	if root == nil {
		return result
	}

	if (root.Left == nil) && (root.Right == nil) {
		result = append(result, []string{root.Val})

		return result
	}

	leftPaths := DepthFirstRecursive(root.Left)

	for _, leftPath := range leftPaths {
		result = append(result, append([]string{root.Val}, leftPath...))
	}

	rightPaths := DepthFirstRecursive(root.Right)

	for _, rightPath := range rightPaths {
		result = append(result, append([]string{root.Val}, rightPath...))
	}

	return result
}
