package treeincludes

func BreadthFirst(root *Node, target string) bool {
	if root == nil {
		return false
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		// get first
		node := queue[0]
		queue = queue[1:]

		if node.Val == target {
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
