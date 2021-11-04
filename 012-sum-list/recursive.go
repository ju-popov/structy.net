package sumlist

func Recursive(head *Node) int {
	if head == nil {
		return 0
	}

	return head.Val + Recursive(head.Next)
}
