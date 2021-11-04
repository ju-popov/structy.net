package isunivaluelist

func Iterative(head *Node) bool {
	for current := head; current != nil; current = current.Next {
		if current.Val != head.Val {
			return false
		}
	}

	return true
}
