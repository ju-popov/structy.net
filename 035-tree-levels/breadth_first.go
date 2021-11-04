package treelevels

type queueValue struct {
	node  *Node
	level int
}

func BreadthFirst(root *Node) [][]string {
	result := make([][]string, 0)

	if root == nil {
		return result
	}

	queue := []queueValue{{node: root, level: 0}}

	for len(queue) > 0 {
		// get first
		node := queue[0].node
		level := queue[0].level
		queue = queue[1:]

		for level >= len(result) {
			result = append(result, make([]string, 0))
		}

		result[level] = append(result[level], node.Val)

		// add last
		if node.Left != nil {
			queue = append(queue, queueValue{node: node.Left, level: level + 1})
		}

		// add last
		if node.Right != nil {
			queue = append(queue, queueValue{node: node.Right, level: level + 1})
		}
	}

	return result
}
