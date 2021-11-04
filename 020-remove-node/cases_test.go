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

func newLinkedList(values ...string) *removenode.Node {
	dummyHead := removenode.NewNode("")
	current := dummyHead

	for _, value := range values {
		current.Next = removenode.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:      "test_00",
		head:      newLinkedList("a", "b", "c", "d", "e", "f"),
		targetVal: "c",
		expected:  newLinkedList("a", "b", "d", "e", "f"),
	},
	{
		name:      "test_01",
		head:      newLinkedList("x", "y", "z"),
		targetVal: "z",
		expected:  newLinkedList("x", "y"),
	},
	{
		name:      "test_02",
		head:      newLinkedList("q", "r", "s"),
		targetVal: "q",
		expected:  newLinkedList("r", "s"),
	},
	{
		name:      "test_03",
		head:      newLinkedList("h", "i", "j", "i"),
		targetVal: "i",
		expected:  newLinkedList("h", "j", "i"),
	},
	{
		name:      "test_04",
		head:      newLinkedList("t"),
		targetVal: "t",
		expected:  nil,
	},
}
