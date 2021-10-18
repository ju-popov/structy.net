package treeincludes

func DepthFirstIterative(root *Node, target string) bool {
	if root == nil {
		return false
	}

	stack := []*Node{root}

	var node *Node

	for len(stack) > 0 {
		// get last
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if node.Value == target {
			return true
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

	return false
}
