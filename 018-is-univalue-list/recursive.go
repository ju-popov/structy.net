package isunivaluelist

func IsUniValue(head *Node, current *Node) bool {
	if (head == nil) || (current == nil) {
		return true
	}

	if head.Val != current.Val {
		return false
	}

	return IsUniValue(head, current.Next)
}

func Recursive(head *Node) bool {
	return IsUniValue(head, head.Next)
}
