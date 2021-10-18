package treevaluecount

func BreadthFirstIterative(root *Node, target int64) int {
	var result int

	if root == nil {
		return result
	}

	queue := []*Node{root}

	var node *Node

	for len(queue) > 0 {
		// get first
		node, queue = queue[0], queue[1:]

		if node.Value == target {
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
