package isunivaluelist_test

import (
	isunivaluelist "github.com/ju-popov/structy.net/018-is-univalue-list"
)

type testCase struct {
	name     string
	head     *isunivaluelist.Node
	expected bool
}

func createLinkedList(values ...interface{}) *isunivaluelist.Node {
	var head *isunivaluelist.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := isunivaluelist.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     createLinkedList(7, 7, 7),
		expected: true,
	},
	{
		name:     "test_01",
		head:     createLinkedList(7, 7, 4),
		expected: false,
	},
	{
		name:     "test_02",
		head:     createLinkedList(2, 2, 2, 2, 2),
		expected: true,
	},
	{
		name:     "test_03",
		head:     createLinkedList(2, 2, 3, 3, 2),
		expected: false,
	},
	{
		name:     "test_04",
		head:     createLinkedList("z"),
		expected: true,
	},
}
