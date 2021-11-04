package treepathfinder

func DepthFirstRecursive(root *Node, target interface{}) []interface{} {
	if root == nil {
		return nil
	}

	result := []interface{}{root.Val}

	if root.Val == target {
		return result
	}

	left := DepthFirstRecursive(root.Left, target)
	if left != nil {
		return append(result, left...)
	}

	right := DepthFirstRecursive(root.Right, target)
	if right != nil {
		return append(result, right...)
	}

	return nil
}
