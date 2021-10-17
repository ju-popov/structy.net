package removenode_test

import (
	removenode "github.com/ju-popov/structy.net/020-remove-node"
)

type testCase struct {
	name      string
	head      *removenode.Node
	targetVal string
	expected  *removenode.Node
}

func createLinkedList(values ...string) *removenode.Node {
	var head *removenode.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := removenode.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:      "test_00",
		head:      createLinkedList("a", "b", "c", "d", "e", "f"),
		targetVal: "c",
		expected:  createLinkedList("a", "b", "d", "e", "f"),
	},
	{
		name:      "test_01",
		head:      createLinkedList("x", "y", "z"),
		targetVal: "z",
		expected:  createLinkedList("x", "y"),
	},
	{
		name:      "test_02",
		head:      createLinkedList("q", "r", "s"),
		targetVal: "q",
		expected:  createLinkedList("r", "s"),
	},
	{
		name:      "test_03",
		head:      createLinkedList("h", "i", "j", "i"),
		targetVal: "i",
		expected:  createLinkedList("h", "j", "i"),
	},
	{
		name:      "test_04",
		head:      createLinkedList("t"),
		targetVal: "t",
		expected:  nil,
	},
}
