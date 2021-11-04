package treelevels

type stackValue struct {
	node  *Node
	level int
}

func DepthFirst(root *Node) [][]string {
	result := make([][]string, 0)

	if root == nil {
		return result
	}

	stack := []stackValue{{node: root, level: 0}}

	for len(stack) > 0 {
		// get last
		node := stack[len(stack)-1].node
		level := stack[len(stack)-1].level
		stack = stack[:len(stack)-1]

		for level >= len(result) {
			result = append(result, make([]string, 0))
		}

		result[level] = append(result[level], node.Val)

		// add last
		if node.Right != nil {
			stack = append(stack, stackValue{node: node.Right, level: level + 1})
		}

		// add last
		if node.Left != nil {
			stack = append(stack, stackValue{node: node.Left, level: level + 1})
		}
	}

	return result
}
