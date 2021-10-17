package treeincludes

func DepthFirstSearch(root *Node, target string) bool {
	if root == nil {
		return false
	}

	if root.Value == target {
		return true
	}

	if DepthFirstSearch(root.Left, target) {
		return true
	}

	if DepthFirstSearch(root.Right, target) {
		return true
	}

	return false
}
