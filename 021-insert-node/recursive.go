package insertnode

func Recursive(head *Node, value string, index int) *Node {
	if index == 0 {
		newNode := NewNode(value)
		newNode.Next = head

		return newNode
	}

	if head != nil {
		head.Next = Recursive(head.Next, value, index-1)
	}

	return head
}
