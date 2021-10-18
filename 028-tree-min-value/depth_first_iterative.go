package treeminvalue

func DepthFirstIterative(root *Node) int64 {
	result := root.Value

	stack := []*Node{root}

	var node *Node

	for len(stack) > 0 {
		// get last
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if node.Value < result {
			result = node.Value
		}

		// add last
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// add last
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
