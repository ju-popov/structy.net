package insertnode_test

import (
	insertnode "github.com/ju-popov/structy.net/021-insert-node"
)

type testCase struct {
	name     string
	head     *insertnode.Node
	value    string
	index    int
	expected *insertnode.Node
}

func createLinkedList(values ...string) *insertnode.Node {
	var head *insertnode.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := insertnode.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     createLinkedList("a", "b", "c", "d"),
		value:    "x",
		index:    2,
		expected: createLinkedList("a", "b", "x", "c", "d"),
	},
	{
		name:     "test_01",
		head:     createLinkedList("a", "b", "c", "d"),
		value:    "v",
		index:    3,
		expected: createLinkedList("a", "b", "c", "v", "d"),
	},
	{
		name:     "test_02",
		head:     createLinkedList("a", "b", "c", "d"),
		value:    "m",
		index:    4,
		expected: createLinkedList("a", "b", "c", "d", "m"),
	},
	{
		name:     "test_03",
		head:     createLinkedList("a", "b"),
		value:    "z",
		index:    0,
		expected: createLinkedList("z", "a", "b"),
	},
}
