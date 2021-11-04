package reverselist_test

import (
	reverselist "github.com/ju-popov/structy.net/015-reverse-list"
)

type testCase struct {
	name     string
	head     *reverselist.Node
	expected *reverselist.Node
}

func newLinkedList(values ...string) *reverselist.Node {
	dummyHead := reverselist.NewNode("")
	current := dummyHead

	for _, value := range values {
		current.Next = reverselist.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList("a", "b", "c", "d", "e", "f"),
		expected: newLinkedList("f", "e", "d", "c", "b", "a"),
	},
	{
		name:     "test_01",
		head:     newLinkedList("x", "y"),
		expected: newLinkedList("y", "x"),
	},
	{
		name:     "test_02",
		head:     newLinkedList("p"),
		expected: newLinkedList("p"),
	},
}
