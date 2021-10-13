package reverselist

type Node struct {
	Value string
	Next  *Node
}

func NewNode(value string) *Node {
	return &Node{Value: value}
}
