package sumlist

func Iterative(head *Node) int {
	var result int

	for current := head; current != nil; current = current.Next {
		result += current.Val
	}

	return result
}
