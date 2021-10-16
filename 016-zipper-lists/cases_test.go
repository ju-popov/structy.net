package zipperlists_test

import (
	zipperlists "github.com/ju-popov/structy.net/016-zipper-lists"
)

type testCase struct {
	name     string
	head1    *zipperlists.Node
	head2    *zipperlists.Node
	expected *zipperlists.Node
}

func createLinkedList(values ...interface{}) *zipperlists.Node {
	var head *zipperlists.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := zipperlists.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head1:    createLinkedList("a", "b", "c"),
		head2:    createLinkedList("x", "y", "z"),
		expected: createLinkedList("a", "x", "b", "y", "c", "z"),
	},
	{
		name:     "test_01",
		head1:    createLinkedList("a", "b", "c", "d", "e", "f"),
		head2:    createLinkedList("x", "y", "z"),
		expected: createLinkedList("a", "x", "b", "y", "c", "z", "d", "e", "f"),
	},
	{
		name:     "test_02",
		head1:    createLinkedList("s", "t"),
		head2:    createLinkedList(1, 2, 3, 4),
		expected: createLinkedList("s", 1, "t", 2, 3, 4),
	},
	{
		name:     "test_03",
		head1:    createLinkedList("w"),
		head2:    createLinkedList(1, 2, 3),
		expected: createLinkedList("w", 1, 2, 3),
	},
	{
		name:     "test_04",
		head1:    createLinkedList(1, 2, 3),
		head2:    createLinkedList("w"),
		expected: createLinkedList(1, "w", 2, 3),
	},
}
