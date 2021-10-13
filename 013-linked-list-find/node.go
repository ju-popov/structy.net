package linkedlistfind

type Node struct {
	value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{value: value}
}
