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

func createLinkedList(values ...string) *getnodevalue.Node {
	var head *getnodevalue.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := getnodevalue.NewNode(values[index])
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
		index:    2,
		expected: "c",
	},
	{
		name:     "test_01",
		head:     createLinkedList("a", "b", "c", "d"),
		index:    3,
		expected: "d",
	},
	{
		name:     "test_02",
		head:     createLinkedList("a", "b", "c", "d"),
		index:    7,
		expected: "",
	},
	{
		name:     "test_03",
		head:     createLinkedList("banana", "mango"),
		index:    0,
		expected: "banana",
	},
	{
		name:     "test_04",
		head:     createLinkedList("banana", "mango"),
		index:    1,
		expected: "mango",
	},
}
