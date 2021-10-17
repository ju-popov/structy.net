package depthfirstvalues

func Iterative(root *Node) []string {
	result := []string{}

	if root == nil {
		return result
	}

	stack := []*Node{root}

	var node *Node
	for len(stack) > 0 {
		// remove last
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]

		result = append(result, node.Value)

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
