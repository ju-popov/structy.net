package linkedlistvalues

type Node struct {
	value string
	Next  *Node
}

func NewNode(value string) *Node {
	return &Node{value: value}
}
