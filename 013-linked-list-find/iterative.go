package linkedlistfind

func Iterative(head *Node, target interface{}) bool {
	for current := head; current != nil; current = current.Next {
		if current.Val == target {
			return true
		}
	}

	return false
}
