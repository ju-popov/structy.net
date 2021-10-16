package mergelists_test

import (
	mergelists "github.com/ju-popov/structy.net/017-merge-lists"
)

type testCase struct {
	name     string
	head1    *mergelists.Node
	head2    *mergelists.Node
	expected *mergelists.Node
}

func createLinkedList(values ...int64) *mergelists.Node {
	var head *mergelists.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := mergelists.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head1:    createLinkedList(5, 7, 10, 12, 20, 28),
		head2:    createLinkedList(6, 8, 9, 25),
		expected: createLinkedList(5, 6, 7, 8, 9, 10, 12, 20, 25, 28),
	},
	{
		name:     "test_01",
		head1:    createLinkedList(5, 7, 10, 12, 20, 28),
		head2:    createLinkedList(1, 8, 9, 10),
		expected: createLinkedList(1, 5, 7, 8, 9, 10, 10, 12, 20, 28),
	},
	{
		name:     "test_02",
		head1:    createLinkedList(30),
		head2:    createLinkedList(15, 67),
		expected: createLinkedList(15, 30, 67),
	},
}
