package lowestcommonancestor

func findPath(root *Node, val string) []string {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return []string{root.Val}
	}

	left := findPath(root.Left, val)
	if left != nil {
		return append(left, root.Val)
	}

	right := findPath(root.Right, val)
	if right != nil {
		return append(right, root.Val)
	}

	return nil
}

func DepthFirstRecursive(root *Node, val1 string, val2 string) string {
	path1 := findPath(root, val1)
	path2 := findPath(root, val2)

	set1 := make(map[string]bool)
	for _, value := range path1 {
		set1[value] = true
	}

	for _, value := range path2 {
		if set1[value] {
			return value
		}
	}

	return ""
}
