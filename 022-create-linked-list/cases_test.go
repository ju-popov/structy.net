package createlinkedlist_test

import (
	createlinkedlist "github.com/ju-popov/structy.net/022-create-linked-list"
)

type testCase struct {
	name     string
	values   []interface{}
	expected *createlinkedlist.Node
}

func newLinkedList(values ...interface{}) *createlinkedlist.Node {
	dummyHead := createlinkedlist.NewNode(nil)
	current := dummyHead

	for _, value := range values {
		current.Next = createlinkedlist.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		values:   []interface{}{"h", "e", "y"},
		expected: newLinkedList("h", "e", "y"),
	},
	{
		name:     "test_01",
		values:   []interface{}{1, 7, 1, 8},
		expected: newLinkedList(1, 7, 1, 8),
	},
	{
		name:     "test_02",
		values:   []interface{}{"a"},
		expected: newLinkedList("a"),
	},
	{
		name:     "test_03",
		values:   []interface{}{},
		expected: nil,
	},
}
