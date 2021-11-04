package leaflist

func DepthFirst(root *Node) []interface{} {
	result := make([]interface{}, 0)

	if root == nil {
		return result
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		// get last
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if (node.Left == nil) && (node.Right == nil) {
			result = append(result, node.Val)
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
