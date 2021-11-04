package reverselist

func helper(head *Node, prev *Node) *Node {
	if head == nil {
		return prev
	}

	next := head.Next
	head.Next = prev

	return helper(next, head)
}

func Recursive(head *Node) *Node {
	return helper(head, nil)
}
