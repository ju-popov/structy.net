package depthfirstvalues

func DepthFirst(root *Node) []string {
	result := []string{}

	if root == nil {
		return result
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		// remove last
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, node.Val)

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
