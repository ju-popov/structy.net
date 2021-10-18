package treevaluecount

func DepthFirstIterative(root *Node, target int64) int {
	var result int

	if root == nil {
		return result
	}

	stack := []*Node{root}

	var node *Node

	for len(stack) > 0 {
		// get last
		node, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if node.Value == target {
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
