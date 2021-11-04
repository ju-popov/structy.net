package treesum

func BreadthFirst(root *Node) int {
	if root == nil {
		return 0
	}

	queue := []*Node{root}

	var result int

	for len(queue) > 0 {
		// get first
		node := queue[0]
		queue = queue[1:]

		result += node.Val

		// put last
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// put last
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
