package depthfirstvalues

func Recursive(root *Node) []string {
	result := []string{}

	if root == nil {
		return result
	}

	result = append(result, root.Value)
	result = append(result, Recursive(root.Left)...)
	result = append(result, Recursive(root.Right)...)

	return result
}
