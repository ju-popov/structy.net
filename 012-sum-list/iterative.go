package sumlist

func Iterative(head *Node) int {
	result := int(0)

	for current := head; current != nil; current = current.Next {
		result += current.Value
	}

	return result
}
