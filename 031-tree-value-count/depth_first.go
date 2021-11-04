package treevaluecount

func DepthFirst(root *Node, target int) int {
	var result int

	if root == nil {
		return result
	}

	stack := []*Node{root}

	for len(stack) > 0 {
		// get last
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node.Val == target {
			result++
		}

		// put last
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

		// put last
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}

	return result
}
