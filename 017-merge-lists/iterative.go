package mergelists

func Iterative(head1 *Node, head2 *Node) *Node {
	dummyHead := NewNode(0)
	current := dummyHead

	for (head1 != nil) && (head2 != nil) {
		if head1.Val <= head2.Val {
			next := head1.Next
			current.Next = head1
			current = head1
			head1 = next
		} else {
			next := head2.Next
			current.Next = head2
			current = head2
			head2 = next
		}
	}

	if head1 != nil {
		current.Next = head1
	}

	if head2 != nil {
		current.Next = head2
	}

	return dummyHead.Next
}
