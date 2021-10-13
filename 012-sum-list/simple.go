package sumlist

func Simple(head *Node) int64 {
	result := int64(0)

	for current := head; current != nil; current = current.Next {
		result += current.value
	}

	return result
}
