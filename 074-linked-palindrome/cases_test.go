package linkedpalindrome_test

import linkedpalindrome "github.com/ju-popov/structy.net/074-linked-palindrome"

type testCase struct {
	name     string
	head     *linkedpalindrome.Node
	expected bool
}

func newLinkedList(values ...int) *linkedpalindrome.Node {
	dummyHead := linkedpalindrome.NewNode(0)
	current := dummyHead

	for _, value := range values {
		current.Next = linkedpalindrome.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList(3, 2, 7, 7, 2, 3),
		expected: true,
	},
	{
		name:     "test_01",
		head:     newLinkedList(3, 2, 4),
		expected: false,
	},
	{
		name:     "test_02",
		head:     newLinkedList(3, 2, 3),
		expected: true,
	},
	{
		name:     "test_03",
		head:     newLinkedList(0, 1, 0, 1, 0),
		expected: true,
	},
	{
		name:     "test_04",
		head:     newLinkedList(0, 1, 0, 1, 1),
		expected: false,
	},
	{
		name:     "test_05",
		head:     newLinkedList(5),
		expected: true,
	},
	{
		name:     "test_06",
		head:     newLinkedList(),
		expected: true,
	},
}
