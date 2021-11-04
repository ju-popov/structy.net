package linkedlistfind_test

import (
	linkedlistfind "github.com/ju-popov/structy.net/013-linked-list-find"
)

type testCase struct {
	name     string
	head     *linkedlistfind.Node
	target   interface{}
	expected bool
}

func newLinkedList(values ...interface{}) *linkedlistfind.Node {
	dummyHead := linkedlistfind.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = linkedlistfind.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList("a", "b", "c", "d"),
		target:   "c",
		expected: true,
	},
	{
		name:     "test_01",
		head:     newLinkedList("a", "b", "c", "d"),
		target:   "d",
		expected: true,
	},
	{
		name:     "test_02",
		head:     newLinkedList("a", "b", "c", "d"),
		target:   "q",
		expected: false,
	},
	{
		name:     "test_03",
		head:     newLinkedList("jason", "leneli"),
		target:   "jason",
		expected: true,
	},
	{
		name:     "test_04",
		head:     newLinkedList(42),
		target:   42,
		expected: true,
	},
	{
		name:     "test_05",
		head:     newLinkedList(42),
		target:   100,
		expected: false,
	},
}
