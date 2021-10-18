package treesum

func DepthFirstIterative(root *Node) int64 {
	if root == nil {
		return 0
	}

	var (
		result int64
		node *Node
	)

	stack := []*Node{root}

	for len(stack) > 0 {
		// get last
		node, stack = stack[len(stack) - 1], stack[:len(stack) - 1]

		result += node.Value

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
