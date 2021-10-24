package sumlist

func Iterative(head *Node) int64 {
	result := int64(0)

	for current := head; current != nil; current = current.Next {
		result += current.Value
	}

	return result
}
