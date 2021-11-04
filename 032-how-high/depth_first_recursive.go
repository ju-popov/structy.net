package howhigh

func maxInt(values ...int) int {
	maxValue := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	return maxValue
}

func DepthFirstRecursive(root *Node) int {
	if root == nil {
		return -1
	}

	left := DepthFirstRecursive(root.Left)
	right := DepthFirstRecursive(root.Right)

	return maxInt(left, right) + 1
}
