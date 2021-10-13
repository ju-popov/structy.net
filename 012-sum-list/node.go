package sumlist

type Node struct {
	value int64
	Next  *Node
}

func NewNode(value int64) *Node {
	return &Node{value: value}
}
