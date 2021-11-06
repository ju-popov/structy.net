package linkedlistcycle

func WithTwoPointers(head *Node) bool {
	slow := head
	fast := head
	firstIteration := true

	for (fast != nil) && (fast.Next != nil) {
		if (fast == slow) && !firstIteration {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
		firstIteration = false
	}

	return false
}
