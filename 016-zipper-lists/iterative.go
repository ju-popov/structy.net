package zipperlists

func Iterative(head1 *Node, head2 *Node) *Node {
	var (
		first *Node
		last  *Node
	)

	for (head1 != nil) || (head2 != nil) {
		if head1 != nil {
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
		}

		if head2 != nil {
			if first == nil {
				first = head2
			}

			if last != nil {
				last.Next = head2
			}

			last = head2

			next := last.Next
			last.Next = nil
			head2 = next
		}
	}

	return first
}
