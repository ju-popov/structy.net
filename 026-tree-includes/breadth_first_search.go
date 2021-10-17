package treeincludes

func BreadthFirstSearch(root *Node, target string) bool {
	if root == nil {
		return false
	}

	queue := []*Node{root}

	var node *Node

	for len(queue) > 0 {
		// get first
		node, queue = queue[0], queue[1:]

		if node.Value == target {
			return true
		}

		// add last
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// add last
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return false
}
