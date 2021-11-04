package mergelists

func Recursive(head1 *Node, head2 *Node) *Node {
	if (head1 == nil) && (head2 == nil) {
		return nil
	}

	if head1 == nil {
		return head2
	}

	if head2 == nil {
		return head1
	}

	if head1.Val <= head2.Val {
		next := head1.Next
		head1.Next = Recursive(next, head2)

		return head1
	}

	next := head2.Next
	head2.Next = Recursive(head1, next)

	return head2
}
