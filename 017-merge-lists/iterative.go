package mergelists

func Iterative(head1 *Node, head2 *Node) *Node {
	var (
		first *Node
		last  *Node
	)

	for (head1 != nil) || (head2 != nil) {
		if (head1 == nil) || ((head2 != nil) && (head2.Value < head1.Value)) {
			head1, head2 = head2, head1
		}

		if first == nil {
			first = head1
		}

		if last != nil {
			last.Next = head1
		}

		last = head1

		next := head1.Next
		last.Next = nil
		head1 = next
	}

	return first
}
