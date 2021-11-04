package removenode

func Iterative(head *Node, targetVal string) *Node {
	var prev *Node

	for current := head; current != nil; current = current.Next {
		if current.Val == targetVal {
			if prev == nil {
				head = head.Next
			} else {
				prev.Next = current.Next
			}

			return head
		}

		prev = current
	}

	return head
}
