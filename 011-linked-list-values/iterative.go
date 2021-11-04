package linkedlistvalues

func Iterative(head *Node) []string {
	result := make([]string, 0)

	for current := head; current != nil; current = current.Next {
		result = append(result, current.Val)
	}

	return result
}
