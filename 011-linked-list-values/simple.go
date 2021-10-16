package linkedlistvalues

func Simple(head *Node) []string {
	result := []string{}

	for current := head; current != nil; current = current.Next {
		result = append(result, current.Value)
	}

	return result
}
