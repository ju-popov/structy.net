package zipperlists

func Iterative(head1 *Node, head2 *Node) *Node {
	var (
		first *Node
		last  *Node
	)

	for (head1 != nil) || (head2 != nil) {
		if head1 == nil {
			head1, head2 = head2, head1
		}

		if first == nil {
			first = head1
		}

		if last != nil {
			last.Next = head1
		}

		last = head1

		next := last.Next
		last.Next = nil
		head1 = next

		head1, head2 = head2, head1
	}

	return first
}
