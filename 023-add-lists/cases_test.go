package addlists_test

import (
	addlists "github.com/ju-popov/structy.net/023-add-lists"
)

type testCase struct {
	name     string
	head1    *addlists.Node
	head2    *addlists.Node
	expected *addlists.Node
}

func createLinkedList(values ...int) *addlists.Node {
	var head *addlists.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := addlists.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head1:    createLinkedList(1, 2, 6),
		head2:    createLinkedList(4, 5, 3),
		expected: createLinkedList(5, 7, 9),
	},
	{
		name:     "test_01",
		head1:    createLinkedList(1, 4, 5, 7),
		head2:    createLinkedList(2, 3),
		expected: createLinkedList(3, 7, 5, 7),
	},
	{
		name:     "test_02",
		head1:    createLinkedList(9, 3),
		head2:    createLinkedList(7, 4),
		expected: createLinkedList(6, 8),
	},
	{
		name:     "test_03",
		head1:    createLinkedList(9, 8),
		head2:    createLinkedList(7, 4),
		expected: createLinkedList(6, 3, 1),
	},
	{
		name:     "test_04",
		head1:    createLinkedList(9, 9, 9),
		head2:    createLinkedList(6),
		expected: createLinkedList(5, 0, 0, 1),
	},
}
