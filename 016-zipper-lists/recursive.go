package zipperlists

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

	next1 := head1.Next
	next2 := head2.Next

	head1.Next = head2
	head2.Next = Recursive(next1, next2)

	return head1
}
