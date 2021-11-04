package treeminvalue

func DepthFirst(root *Node) int64 {
	result := root.Val

	stack := []*Node{root}

	for len(stack) > 0 {
		// get last
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Val < result {
			result = node.Val
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
