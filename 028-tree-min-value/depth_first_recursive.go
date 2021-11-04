package treeminvalue

func DepthFirstRecursive(root *Node) int64 {
	result := root.Val

	if root.Left != nil {
		if resultLeft := DepthFirstRecursive(root.Left); resultLeft < result {
			result = resultLeft
		}
	}

	if root.Right != nil {
		if resultRight := DepthFirstRecursive(root.Right); resultRight < result {
			result = resultRight
		}
	}

	return result
}
