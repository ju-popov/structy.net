package sumlist_test

import (
	sumlist "github.com/ju-popov/structy.net/012-sum-list"
)

type testCase struct {
	name     string
	head     *sumlist.Node
	expected int
}

func newLinkedList(values ...int) *sumlist.Node {
	dummyHead := sumlist.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = sumlist.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList(2, 8, 3, -1, 7),
		expected: 19,
	},
	{
		name:     "test_01",
		head:     newLinkedList(38, 4),
		expected: 42,
	},
	{
		name:     "test_02",
		head:     newLinkedList(100),
		expected: 100,
	},
	{
		name:     "test_03",
		head:     newLinkedList(),
		expected: 0,
	},
}
