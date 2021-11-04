package addlists

func addLists(head1 *Node, head2 *Node, carry int) *Node {
	if (head1 == nil) && (head2 == nil) && (carry == 0) {
		return nil
	}

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

	newNode.Next = addLists(head1, head2, carry)

	return newNode
}

func Recursive(head1 *Node, head2 *Node) *Node {
	return addLists(head1, head2, 0)
}
