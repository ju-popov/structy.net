package treeminvalue

func BreadthFirst(root *Node) int {
	result := root.Val

	queue := []*Node{root}

	for len(queue) > 0 {
		// get first
		node := queue[0]
		queue = queue[1:]

		if node.Val < result {
			result = node.Val
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

	return result
}
