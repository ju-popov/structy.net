package linkedpalindrome

func Iterative(head *Node) bool {
	values := make([]interface{}, 0)

	for current := head; current != nil; current = current.Next {
		values = append(values, current.Val)
	}

	for index := 0; index < len(values)/2; index++ {
		if values[index] != values[len(values)-1-index] {
			return false
		}
	}

	return true
}
