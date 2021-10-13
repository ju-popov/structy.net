package linkedlistfind

func Simple(head *Node, target interface{}) bool {
	for current := head; current != nil; current = current.Next {
		if current.value == target {
			return true
		}
	}

	return false
}
