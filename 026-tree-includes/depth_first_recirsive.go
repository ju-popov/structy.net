package treeincludes

func DepthFirstRecursive(root *Node, target string) bool {
	if root == nil {
		return false
	}

	if root.Value == target {
		return true
	}

	if DepthFirstRecursive(root.Left, target) {
		return true
	}

	if DepthFirstRecursive(root.Right, target) {
		return true
	}

	return false
}
