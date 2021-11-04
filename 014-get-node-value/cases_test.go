package getnodevalue_test

import (
	getnodevalue "github.com/ju-popov/structy.net/014-get-node-value"
)

type testCase struct {
	name     string
	head     *getnodevalue.Node
	index    int
	expected string
}

func newLinkedList(values ...string) *getnodevalue.Node {
	dummyHead := getnodevalue.NewNode("")
	current := dummyHead

	for _, value := range values {
		current.Next = getnodevalue.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList("a", "b", "c", "d"),
		index:    2,
		expected: "c",
	},
	{
		name:     "test_01",
		head:     newLinkedList("a", "b", "c", "d"),
		index:    3,
		expected: "d",
	},
	{
		name:     "test_02",
		head:     newLinkedList("a", "b", "c", "d"),
		index:    7,
		expected: "",
	},
	{
		name:     "test_03",
		head:     newLinkedList("banana", "mango"),
		index:    0,
		expected: "banana",
	},
	{
		name:     "test_04",
		head:     newLinkedList("banana", "mango"),
		index:    1,
		expected: "mango",
	},
}
