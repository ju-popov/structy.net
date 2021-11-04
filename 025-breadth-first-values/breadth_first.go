package breadthfirstvalues

func BreadthFirst(root *Node) []string {
	result := make([]string, 0)

	if root == nil {
		return result
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		// get first
		node := queue[0]
		queue = queue[1:]

		result = append(result, node.Val)

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
