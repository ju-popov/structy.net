package addlists

func Iterative(head1 *Node, head2 *Node) *Node {
	dummyHead := NewNode(0)
	current := dummyHead
	carry := 0

	for (head1 != nil) || (head2 != nil) || (carry > 0) {
		val1 := 0
		if head1 != nil {
			val1 = head1.Val
		}

		val2 := 0
		if head2 != nil {
			val2 = head2.Val
		}

		total := val1 + val2 + carry

		//nolint:gomnd
		newNode := NewNode(total % 10)
		//nolint:gomnd
		carry = total / 10

		if head1 != nil {
			head1 = head1.Next
		}

		if head2 != nil {
			head2 = head2.Next
		}

		current.Next = newNode
		current = newNode
	}

	return dummyHead.Next
}
