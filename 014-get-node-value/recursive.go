package getnodevalue

func Recursive(head *Node, index int) string {
	if head == nil {
		return ""
	}

	if index == 0 {
		return head.Val
	}

	return Recursive(head.Next, index-1)
}
