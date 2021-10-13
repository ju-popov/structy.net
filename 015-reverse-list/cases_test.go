package reverselist_test

import (
	reverselist "github.com/ju-popov/structy.net/015-reverse-list"
)

type testCase struct {
	name     string
	head     *reverselist.Node
	expected *reverselist.Node
}

func createLinkedList(values ...string) *reverselist.Node {
	var head *reverselist.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := reverselist.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     createLinkedList("a", "b", "c", "d", "e", "f"),
		expected: createLinkedList("f", "e", "d", "c", "b", "a"),
	},
	{
		name:     "test_01",
		head:     createLinkedList("x", "y"),
		expected: createLinkedList("y", "x"),
	},
	{
		name:     "test_02",
		head:     createLinkedList("p"),
		expected: createLinkedList("p"),
	},
}
