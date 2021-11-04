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
	var (
		first *getnodevalue.Node
		last  *getnodevalue.Node
	)

	for _, value := range values {
		node := getnodevalue.NewNode(value)

		if first == nil {
			first = node
		}

		if last != nil {
			last.Next = node
		}

		last = node
	}

	return first
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
