package treesum

func BreadthFirstIterative(root *Node) int64 {
	if root == nil {
		return 0
	}

	queue := []*Node{root}

	var (
		node   *Node
		result int64
	)

	for len(queue) > 0 {
		// get first
		node, queue = queue[0], queue[1:]

		result += node.Value

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
