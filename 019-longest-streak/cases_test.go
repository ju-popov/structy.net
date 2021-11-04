package longeststreak_test

import (
	longeststreak "github.com/ju-popov/structy.net/019-longest-streak"
)

type testCase struct {
	name     string
	head     *longeststreak.Node
	expected int
}

func createLinkedList(values ...int) *longeststreak.Node {
	var head *longeststreak.Node

	for index := len(values) - 1; index >= 0; index-- {
		node := longeststreak.NewNode(values[index])
		node.Next = head
		head = node
	}

	return head
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     createLinkedList(5, 5, 7, 7, 7, 6),
		expected: 3,
	},
	{
		name:     "test_01",
		head:     createLinkedList(3, 3, 3, 3, 9, 9),
		expected: 4,
	},
	{
		name:     "test_02",
		head:     createLinkedList(9, 9, 1, 9, 9, 9),
		expected: 3,
	},
	{
		name:     "test_03",
		head:     createLinkedList(5, 5),
		expected: 2,
	},
	{
		name:     "test_04",
		head:     createLinkedList(4),
		expected: 1,
	},
	{
		name:     "test_05",
		head:     nil,
		expected: 0,
	},
}
