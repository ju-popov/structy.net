package removenode

func Iterative(head *Node, targetVal string) *Node {
	var previous *Node

	for current := head; current != nil; current = current.Next {
		if current.Value == targetVal {
			if previous == nil {
				head = head.Next
			} else {
				previous.Next = current.Next
			}

			return head
		}

		previous = current
	}

	return head
}
