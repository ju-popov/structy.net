package createlinkedlist

func Recursive(values []interface{}) *Node {
	if len(values) <= 0 {
		return nil
	}

	newNode := NewNode(values[0])

	newNode.Next = Recursive(values[1:])

	return newNode
}
