package linkedlistcycle

func WithSet(head *Node) bool {
	visited := make(map[*Node]bool)

	for current := head; current != nil; current = current.Next {
		if visited[current] {
			return true
		}

		visited[current] = true
	}

	return false
}
