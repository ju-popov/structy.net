package insertnode

func Iterative(head *Node, value string, index int) *Node {
	//nolint:varnamelen
	i := 0

	for current, prev := head, (*Node)(nil); (current != nil) || (prev != nil); current, prev = current.Next, current {
		if i == index {
			newNode := NewNode(value)
			newNode.Next = current

			if prev == nil {
				head = newNode
			} else {
				prev.Next = newNode
			}

			break
		}

		i++
	}

	return head
}
