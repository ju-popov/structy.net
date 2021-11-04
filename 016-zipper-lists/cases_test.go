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

func newLinkedList(values ...interface{}) *zipperlists.Node {
	dummyHead := zipperlists.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = zipperlists.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head1:    newLinkedList("a", "b", "c"),
		head2:    newLinkedList("x", "y", "z"),
		expected: newLinkedList("a", "x", "b", "y", "c", "z"),
	},
	{
		name:     "test_01",
		head1:    newLinkedList("a", "b", "c", "d", "e", "f"),
		head2:    newLinkedList("x", "y", "z"),
		expected: newLinkedList("a", "x", "b", "y", "c", "z", "d", "e", "f"),
	},
	{
		name:     "test_02",
		head1:    newLinkedList("s", "t"),
		head2:    newLinkedList(1, 2, 3, 4),
		expected: newLinkedList("s", 1, "t", 2, 3, 4),
	},
	{
		name:     "test_03",
		head1:    newLinkedList("w"),
		head2:    newLinkedList(1, 2, 3),
		expected: newLinkedList("w", 1, 2, 3),
	},
	{
		name:     "test_04",
		head1:    newLinkedList(1, 2, 3),
		head2:    newLinkedList("w"),
		expected: newLinkedList(1, "w", 2, 3),
	},
}
