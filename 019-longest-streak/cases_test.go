package longeststreak_test

import (
	longeststreak "github.com/ju-popov/structy.net/019-longest-streak"
)

type testCase struct {
	name     string
	head     *longeststreak.Node
	expected int
}

func newLinkedList(values ...int) *longeststreak.Node {
	dummyHead := longeststreak.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = longeststreak.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList(5, 5, 7, 7, 7, 6),
		expected: 3,
	},
	{
		name:     "test_01",
		head:     newLinkedList(3, 3, 3, 3, 9, 9),
		expected: 4,
	},
	{
		name:     "test_02",
		head:     newLinkedList(9, 9, 1, 9, 9, 9),
		expected: 3,
	},
	{
		name:     "test_03",
		head:     newLinkedList(5, 5),
		expected: 2,
	},
	{
		name:     "test_04",
		head:     newLinkedList(4),
		expected: 1,
	},
	{
		name:     "test_05",
		head:     nil,
		expected: 0,
	},
}
