package zipperlists

func AddNode(current *Node, head1 *Node, head2 *Node) *Node {
	if head1 == nil {
		if head2 == nil {
			return current
		}

		head1, head2 = head2, head1
	}

	next := head1.Next
	head1.Next = nil

	if current != nil {
		current.Next = head1
	}

	current = head1

	AddNode(current, head2, next)

	return current
}

func Recursive(head1 *Node, head2 *Node) *Node {
	return AddNode(nil, head1, head2)
}
