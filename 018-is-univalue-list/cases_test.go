package isunivaluelist_test

import (
	isunivaluelist "github.com/ju-popov/structy.net/018-is-univalue-list"
)

type testCase struct {
	name     string
	head     *isunivaluelist.Node
	expected bool
}

func newLinkedList(values ...interface{}) *isunivaluelist.Node {
	dummyHead := isunivaluelist.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = isunivaluelist.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList(7, 7, 7),
		expected: true,
	},
	{
		name:     "test_01",
		head:     newLinkedList(7, 7, 4),
		expected: false,
	},
	{
		name:     "test_02",
		head:     newLinkedList(2, 2, 2, 2, 2),
		expected: true,
	},
	{
		name:     "test_03",
		head:     newLinkedList(2, 2, 3, 3, 2),
		expected: false,
	},
	{
		name:     "test_04",
		head:     newLinkedList("z"),
		expected: true,
	},
}
