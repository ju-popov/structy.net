package leaflist

func DepthFirstIterative(root *Node) []interface{} {
	result := []interface{}{}

	if root == nil {
		return result
	}

	stack := []*Node{root}

	var node *Node

	for len(stack) > 0 {
		// get last
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if (node.Left == nil) && (node.Right == nil) {
			result = append(result, node.Value)
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
