package createlinkedlist_test

import (
	createlinkedlist "github.com/ju-popov/structy.net/022-create-linked-list"
)

type testCase struct {
	name     string
	values   []interface{}
	expected *createlinkedlist.Node
}

func createLinkedList(values ...interface{}) *createlinkedlist.Node {
	var head *createlinkedlist.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := createlinkedlist.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		values:   []interface{}{"h", "e", "y"},
		expected: createLinkedList("h", "e", "y"),
	},
	{
		name:     "test_01",
		values:   []interface{}{1, 7, 1, 8},
		expected: createLinkedList(1, 7, 1, 8),
	},
	{
		name:     "test_02",
		values:   []interface{}{"a"},
		expected: createLinkedList("a"),
	},
	{
		name:     "test_03",
		values:   []interface{}{},
		expected: nil,
	},
}
