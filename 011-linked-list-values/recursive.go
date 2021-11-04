package linkedlistvalues

func helper(head *Node, result []string) []string {
	if head == nil {
		return result
	}

	result = append(result, head.Val)

	return helper(head.Next, result)
}

func Recursive(head *Node) []string {
	return helper(head, make([]string, 0))
}
