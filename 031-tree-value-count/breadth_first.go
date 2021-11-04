package treevaluecount

func BreadthFirst(root *Node, target int) int {
	var result int

	if root == nil {
		return result
	}

	queue := []*Node{root}

	for len(queue) > 0 {
		// get first
		node := queue[0]
		queue = queue[1:]

		if node.Val == target {
			result++
		}

		// put last
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// put last
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return result
}
