package getnodevalue

func Simple(head *Node, index int64) string {
	for current, i := head, int64(0); current != nil; current, i = current.Next, i+1 {
		if i == index {
			return current.Value
		}
	}

	return ""
}
