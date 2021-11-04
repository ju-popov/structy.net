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

func newLinkedList(values ...int) *addlists.Node {
	dummyHead := addlists.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = addlists.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head1:    newLinkedList(1, 2, 6),
		head2:    newLinkedList(4, 5, 3),
		expected: newLinkedList(5, 7, 9),
	},
	{
		name:     "test_01",
		head1:    newLinkedList(1, 4, 5, 7),
		head2:    newLinkedList(2, 3),
		expected: newLinkedList(3, 7, 5, 7),
	},
	{
		name:     "test_02",
		head1:    newLinkedList(9, 3),
		head2:    newLinkedList(7, 4),
		expected: newLinkedList(6, 8),
	},
	{
		name:     "test_03",
		head1:    newLinkedList(9, 8),
		head2:    newLinkedList(7, 4),
		expected: newLinkedList(6, 3, 1),
	},
	{
		name:     "test_04",
		head1:    newLinkedList(9, 9, 9),
		head2:    newLinkedList(6),
		expected: newLinkedList(5, 0, 0, 1),
	},
}
