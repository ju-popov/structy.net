package middlevalue

func WithArray(head *Node) string {
	values := make([]string, 0)

	for current := head; current != nil; current = current.Next {
		values = append(values, current.Val)
	}

	return values[len(values)/2]
}
