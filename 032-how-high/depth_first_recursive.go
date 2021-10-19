package howhigh

func helper(root *Node) int {
	if root == nil {
		return 0
	}

	var (
		left  int
		right int
	)

	if left, right = helper(root.Left), helper(root.Right); left > right {
		return left + 1
	}

	return right + 1
}

func DepthFirstRecursive(root *Node) int {
	return helper(root) - 1
}
