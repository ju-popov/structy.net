package isunivaluelist

func IsUniValue(first *Node, current *Node) bool {
	if current == nil {
		return true
	}

	if first.Value != current.Value {
		return false
	}

	return IsUniValue(first, current.Next)
}

func Recursive(head *Node) bool {
	return IsUniValue(head, head.Next)
}
