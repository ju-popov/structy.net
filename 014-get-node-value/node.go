package getnodevalue

import (
	"bytes"
	"fmt"
)

type Node struct {
	Val  string
	Next *Node
}

func NewNode(value string) *Node {
	return &Node{Val: value}
}

func (head *Node) String() string {
	buf := bytes.NewBufferString("")

	for current := head; current != nil; current = current.Next {
		if current != head {
			if _, err := fmt.Fprint(buf, " -> "); err != nil {
				panic(err)
			}
		}

		if _, err := fmt.Fprintf(buf, "%v", current.Val); err != nil {
			panic(err)
		}
	}

	return buf.String()
}

func (head *Node) Copy() *Node {
	dummyHead := NewNode("")
	last := dummyHead

	for current := head; current != nil; current = current.Next {
		last.Next = NewNode(current.Val)
		last = last.Next
	}

	return dummyHead.Next
}
