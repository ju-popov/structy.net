package isunivaluelist

func Iterative(head *Node) bool {
	var first *Node

	for current := head; current != nil; current = current.Next {
		if first == nil {
			first = current
		} else if first.Value != current.Value {
			return false
		}
	}

	return true
}
