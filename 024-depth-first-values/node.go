package depthfirstvalues

type Node struct {
	Value string
	Left  *Node
	Right *Node
}

func NewNode(value string) *Node {
	return &Node{Value: value}
}
