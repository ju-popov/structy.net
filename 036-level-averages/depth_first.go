package levelaverages

type stackValue struct {
	node  *Node
	level int
}

func depthFirstFillLevels(root *Node) [][]float64 {
	result := make([][]float64, 0)

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
			result = append(result, make([]float64, 0))
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

func DepthFirst(root *Node) []float64 {
	levels := depthFirstFillLevels(root)

	result := make([]float64, 0)

	for _, level := range levels {
		var sum float64
		for _, value := range level {
			sum += value
		}

		average := sum / float64(len(level))
		result = append(result, average)
	}

	return result
}
