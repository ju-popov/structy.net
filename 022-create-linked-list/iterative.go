package createlinkedlist

func Iterative(values []interface{}) *Node {
	dummyHead := NewNode(nil)
	current := dummyHead

	for _, value := range values {
		current.Next = NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}
