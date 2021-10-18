package maxroottoleafpathsum

func DepthFirstRecursive(root *Node) int64 {
	if root.Left != nil {
		if root.Right != nil {
			// left != nil, right != nil
			var (
				left  int64
				right int64
			)

			if left, right = DepthFirstRecursive(root.Left), DepthFirstRecursive(root.Right); left > right {
				// left != nil, right != nil, left > right
				return root.Value + left
			}

			// left != nil, right != nil, left <= right
			return root.Value + right
		}

		// left != nil, right == nil
		return root.Value + DepthFirstRecursive(root.Left)
	}

	if root.Right != nil {
		// left == nil, right != nil
		return root.Value + DepthFirstRecursive(root.Right)
	}

	// left == nil, right == nil
	return root.Value
}
