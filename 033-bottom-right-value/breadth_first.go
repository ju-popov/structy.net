package bottomrightvalue

func BreadthFirst(root *Node) interface{} {
	queue := []*Node{root}

	var node *Node

	for len(queue) > 0 {
		// get first
		node = queue[0]
		queue = queue[1:]

		// add last
		if node.Left != nil {
			queue = append(queue, node.Left)
		}

		// add last
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}

	return node.Val
}
