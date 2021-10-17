package createlinkedlist

func Iterative(values []interface{}) *Node {
	var (
		head     *Node
		previous *Node
	)

	for _, value := range values {
		newNode := NewNode(value)

		if head == nil {
			head = newNode
		}

		if previous != nil {
			previous.Next = newNode
		}

		previous = newNode
	}

	return head
}
