package bottomrightvalue

func helper(root *Node, level int, results map[int]interface{}) interface{} {
	if root == nil {
		return nil
	}

	if _, ok := results[level]; !ok {
		results[level] = root.Val
	}

	helper(root.Right, level+1, results)
	helper(root.Left, level+1, results)

	return results[len(results)-1]
}

func DepthFirstRecusrsive(root *Node) interface{} {
	return helper(root, 0, make(map[int]interface{}))
}
