package linkedlistfind

func Recursive(head *Node, target interface{}) bool {
	if head == nil {
		return false
	}

	if head.Val == target {
		return true
	}

	return Recursive(head.Next, target)
}
