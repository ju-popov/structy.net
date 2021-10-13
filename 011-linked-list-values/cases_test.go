package linkedlistvalues_test

import linkedlistvalues "github.com/ju-popov/structy.net/011-linked-list-values"

type testCase struct {
	name     string
	head     *linkedlistvalues.Node
	expected []string
}

func createLinkedList(values ...string) *linkedlistvalues.Node {
	var head *linkedlistvalues.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := linkedlistvalues.NewNode(values[index])
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
		expected: []string{"a", "b", "c", "d"},
	},
	{
		name:     "test_01",
		head:     createLinkedList("x", "y"),
		expected: []string{"x", "y"},
	},
	{
		name:     "test_02",
		head:     createLinkedList("q"),
		expected: []string{"q"},
	},
	{
		name:     "test_03",
		head:     createLinkedList(),
		expected: []string{},
	},
}
