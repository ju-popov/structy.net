package zipperlists

func Iterative(head1 *Node, head2 *Node) *Node {
	dummyHead := NewNode(nil)
	current := dummyHead

	for (head1 != nil) && (head2 != nil) {
		next1 := head1.Next
		next2 := head2.Next

		current.Next = head1
		head1.Next = head2
		current = head2

		head1 = next1
		head2 = next2
	}

	if head1 != nil {
		if current != nil {
			current.Next = head1
		}
	}

	if head2 != nil {
		if current != nil {
			current.Next = head2
		}
	}

	return dummyHead.Next
}
