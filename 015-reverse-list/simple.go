package reverselist

func Simple(head *Node) *Node {
	var prev *Node

	for current := head; current != nil; {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	return prev
}
