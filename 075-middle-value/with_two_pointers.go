package middlevalue

func WithTwoPointers(head *Node) string {
	slow := head
	fast := head

	for (fast != nil) && (fast.Next != nil) {
		slow = slow.Next
		fast = fast.Next.Next
	}

	return slow.Val
}
