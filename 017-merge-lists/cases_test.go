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

func newLinkedList(values ...int) *mergelists.Node {
	dummyHead := mergelists.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = mergelists.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head1:    newLinkedList(5, 7, 10, 12, 20, 28),
		head2:    newLinkedList(6, 8, 9, 25),
		expected: newLinkedList(5, 6, 7, 8, 9, 10, 12, 20, 25, 28),
	},
	{
		name:     "test_01",
		head1:    newLinkedList(5, 7, 10, 12, 20, 28),
		head2:    newLinkedList(1, 8, 9, 10),
		expected: newLinkedList(1, 5, 7, 8, 9, 10, 10, 12, 20, 28),
	},
	{
		name:     "test_02",
		head1:    newLinkedList(30),
		head2:    newLinkedList(15, 67),
		expected: newLinkedList(15, 30, 67),
	},
}
