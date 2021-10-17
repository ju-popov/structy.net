package depthfirstvalues

func Iterative(root *Node) []string {
	result := []string{}

	if root == nil {
		return result
	}

	stack := []*Node{root}

	var node *Node
	for len(stack) > 0 {
		// get first
		node, stack = stack[0], stack[1:]

		result = append(result, node.Value)

		// add first
		if node.Right != nil {
			stack = append([]*Node{node.Right}, stack...)
		}

		// add first
		if node.Left != nil {
			stack = append([]*Node{node.Left}, stack...)
		}
	}

	return result
}
