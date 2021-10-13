package sumlist_test

import (
	sumlist "github.com/ju-popov/structy.net/012-sum-list"
)

type testCase struct {
	name     string
	head     *sumlist.Node
	expected int64
}

func createLinkedList(values ...int64) *sumlist.Node {
	var head *sumlist.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := sumlist.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     createLinkedList(2, 8, 3, -1, 7),
		expected: 19,
	},
	{
		name:     "test_01",
		head:     createLinkedList(38, 4),
		expected: 42,
	},
	{
		name:     "test_02",
		head:     createLinkedList(100),
		expected: 100,
	},
	{
		name:     "test_03",
		head:     createLinkedList(),
		expected: 0,
	},
}
