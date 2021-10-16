package zipperlists

func AddNode(head *Node, head1 *Node, head2 *Node) *Node {
	var next *Node
	var current *Node
	if head1 != nil {
		next = head1.Next
		head1.Next = nil
		if head == nil {
			head = head1
			current = head
		} else {
			head.Next = head1
			current = head1
		}
	}

	if head2 != nil {
		AddNode(current, head2, next)
	} else if next != nil {
		AddNode(current, next, head2)
	}

	return head
}

func Recursive(head1 *Node, head2 *Node) *Node {
	return AddNode(nil, head1, head2)
}
