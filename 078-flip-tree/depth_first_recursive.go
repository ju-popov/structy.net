package fliptree

func DepthFirstRecursive(root *Node) *Node {
	if root == nil {
		return nil
	}

	DepthFirstRecursive(root.Left)
	DepthFirstRecursive(root.Right)

	root.Left, root.Right = root.Right, root.Left

	return root
}
