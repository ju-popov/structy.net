package levelaverages

func depthFirstRecursiveFillLevels(root *Node, level int, result *[][]float64) {
	if root == nil {
		return
	}

	for len(*result) <= level {
		*result = append(*result, []float64{})
	}

	(*result)[level] = append((*result)[level], root.Value)

	if root.Left != nil {
		depthFirstRecursiveFillLevels(root.Left, level + 1, result)
	}

	if root.Right != nil {
		depthFirstRecursiveFillLevels(root.Right, level + 1, result)
	}
}

func DepthFirstRecursive(root *Node) []float64 {
	levels := [][]float64{}

	depthFirstRecursiveFillLevels(root, 0, &levels)

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
