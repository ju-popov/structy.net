package maxroottoleafpathsum

func maxInt(values ...int) int {
	maxValue := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	return maxValue
}

func DepthFirstRecursive(root *Node) int {
	if root.Left != nil {
		if root.Right != nil {
			// left != nil, right != nil
			return root.Val + maxInt(DepthFirstRecursive(root.Left), DepthFirstRecursive(root.Right))
		}

		// left != nil, right == nil
		return root.Val + DepthFirstRecursive(root.Left)
	}

	if root.Right != nil {
		// left == nil, right != nil
		return root.Val + DepthFirstRecursive(root.Right)
	}

	// left == nil, right == nil
	return root.Val
}
