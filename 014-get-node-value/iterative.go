package getnodevalue

func Iterative(head *Node, index int) string {
	for current, i := head, int(0); current != nil; current, i = current.Next, i+1 {
		if i == index {
			return current.Val
		}
	}

	return ""
}
