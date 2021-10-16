package zipperlists

import (
	"bytes"
	"fmt"
)

type Node struct {
	Value interface{}
	Next  *Node
}

func NewNode(value interface{}) *Node {
	return &Node{Value: value}
}

func (head *Node) String() string {
	buf := bytes.NewBufferString("")

	for current := head; current != nil; current = current.Next {
		if current != head {
			if _, err := fmt.Fprint(buf, " -> "); err != nil {
				panic(err)
			}
		}
		if _, err := fmt.Fprintf(buf, "%v", current.Value); err != nil {
			panic(err)
		}
	}

	return buf.String()
}

func (head *Node) Copy() *Node {
	var (
		last *Node
		first *Node
	)

	for current := head; current != nil; current = current.Next {
		n := NewNode(current.Value)
		if last != nil {
			last.Next = n
		}
		last = n
		if first == nil {
			first = last
		}
	}

	return first
}
