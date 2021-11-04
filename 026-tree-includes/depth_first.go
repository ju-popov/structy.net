package treeincludes

func DepthFirst(root *Node, target string) bool {
	if root == nil {
		return false
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		// get last
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Val == target {
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
