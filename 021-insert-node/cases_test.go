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

func newLinkedList(values ...string) *insertnode.Node {
	dummyHead := insertnode.NewNode("")
	current := dummyHead

	for _, value := range values {
		current.Next = insertnode.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList("a", "b", "c", "d"),
		value:    "x",
		index:    2,
		expected: newLinkedList("a", "b", "x", "c", "d"),
	},
	{
		name:     "test_01",
		head:     newLinkedList("a", "b", "c", "d"),
		value:    "v",
		index:    3,
		expected: newLinkedList("a", "b", "c", "v", "d"),
	},
	{
		name:     "test_02",
		head:     newLinkedList("a", "b", "c", "d"),
		value:    "m",
		index:    4,
		expected: newLinkedList("a", "b", "c", "d", "m"),
	},
	{
		name:     "test_03",
		head:     newLinkedList("a", "b"),
		value:    "z",
		index:    0,
		expected: newLinkedList("z", "a", "b"),
	},
}
