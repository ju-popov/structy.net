package linkedlistcycle

func helper(head *Node, visited map[*Node]bool) bool {
	if head == nil {
		return false
	}

	if visited[head] {
		return true
	}

	visited[head] = true

	return helper(head.Next, visited)
}

func WithSetRecursive(head *Node) bool {
	visited := make(map[*Node]bool)

	return helper(head, visited)
}
