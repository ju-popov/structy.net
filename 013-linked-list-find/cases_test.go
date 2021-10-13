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

func createLinkedList(values ...interface{}) *linkedlistfind.Node {
	var head *linkedlistfind.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := linkedlistfind.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     createLinkedList("a", "b", "c", "d"),
		target:   "c",
		expected: true,
	},
	{
		name:     "test_01",
		head:     createLinkedList("a", "b", "c", "d"),
		target:   "d",
		expected: true,
	},
	{
		name:     "test_02",
		head:     createLinkedList("a", "b", "c", "d"),
		target:   "q",
		expected: false,
	},
	{
		name:     "test_03",
		head:     createLinkedList("jason", "leneli"),
		target:   "jason",
		expected: true,
	},
	{
		name:     "test_04",
		head:     createLinkedList(42),
		target:   42,
		expected: true,
	},
	{
		name:     "test_05",
		head:     createLinkedList(42),
		target:   100,
		expected: false,
	},
}
