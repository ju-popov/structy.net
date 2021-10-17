package removenode

func Recursive(head *Node, targetVal string) *Node {
	if head == nil {
		return nil
	}

	if head.Value == targetVal {
		return head.Next
	}

	head.Next = Recursive(head.Next, targetVal)

	return head
}
