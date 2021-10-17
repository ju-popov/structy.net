package insertnode

func Iterative(head *Node, value string, index int) *Node {
	for i, current, previous := 0, head, (*Node)(nil); (current != nil) || (previous != nil); i, current, previous = i+1, current.Next, current {
		if i == index {
			newNode := NewNode(value)
			newNode.Next = current

			if previous == nil {
				head = newNode
			} else {
				previous.Next = newNode
			}

			return head
		}

		if current == nil {
			return head
		}
	}

	return head
}
