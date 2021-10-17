package breadthfirstvalues

func Iterative(root *Node) []string {
	result := []string{}

	if root == nil {
		return result
	}

	queue := []*Node{root}

	var node *Node
	for len(queue) > 0 {
		// get first
		node, queue = queue[0], queue[1:]

		result = append(result, node.Value)

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
