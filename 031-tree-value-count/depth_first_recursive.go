package treevaluecount

func DepthFirstRecursive(root *Node, target int) int {
	var result int

	if root == nil {
		return result
	}

	if root.Val == target {
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
