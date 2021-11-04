package leaflist

func DepthFirstRecursive(root *Node) []interface{} {
	result := make([]interface{}, 0)

	if root == nil {
		return result
	}

	if (root.Left == nil) && (root.Right == nil) {
		return []interface{}{root.Val}
	}

	if root.Left != nil {
		result = append(result, DepthFirstRecursive(root.Left)...)
	}

	if root.Right != nil {
		result = append(result, DepthFirstRecursive(root.Right)...)
	}

	return result
}
