package levelaverages

func depthFirstRecursiveFillLevels(root *Node, level int, result [][]float64) [][]float64 {
	if root == nil {
		return result
	}

	for level >= len(result) {
		result = append(result, make([]float64, 0))
	}

	result[level] = append(result[level], root.Val)

	if root.Left != nil {
		result = depthFirstRecursiveFillLevels(root.Left, level+1, result)
	}

	if root.Right != nil {
		result = depthFirstRecursiveFillLevels(root.Right, level+1, result)
	}

	return result
}

func DepthFirstRecursive(root *Node) []float64 {
	levels := depthFirstRecursiveFillLevels(root, 0, make([][]float64, 0))

	result := []float64{}

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
