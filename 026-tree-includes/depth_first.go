package treeincludes

func DepthFirst(root *Node, target string) bool {
	if root == nil {
		return false
	}

	if root.Value == target {
		return true
	}

	if DepthFirst(root.Left, target) {
		return true
	}

	if DepthFirst(root.Right, target) {
		return true
	}

	return false
}
