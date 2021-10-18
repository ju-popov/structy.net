package treepathfinder

func depthFirstRecursiveHelper(root *Node, target interface{}, path []interface{}) []interface{} {
	if root == nil {
		return nil
	}

	if root.Value == target {
		return append(path, root.Value)
	}

	if root.Left != nil {
		if result := depthFirstRecursiveHelper(root.Left, target, append(path, root.Value)); result != nil {
			return result
		}
	}

	if root.Right != nil {
		if result := depthFirstRecursiveHelper(root.Right, target, append(path, root.Value)); result != nil {
			return result
		}
	}

	return nil
}

func DepthFirstRecursive(root *Node, target interface{}) []interface{} {
	return depthFirstRecursiveHelper(root, target, []interface{}{})
}
