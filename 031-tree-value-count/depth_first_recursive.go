package treevaluecount

func DepthFirstRecursive(root *Node, target int64) int {
	var result int

	if root == nil {
		return result
	}

	if root.Value == target {
		result++
	}

	if root.Left != nil {
		result += DepthFirstRecursive(root.Left, target)
	}

	if root.Right != nil {
		result += DepthFirstRecursive(root.Right, target)
	}

	return result
}
