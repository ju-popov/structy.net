package linkedlistvalues_test

import linkedlistvalues "github.com/ju-popov/structy.net/011-linked-list-values"

type testCase struct {
	name     string
	head     *linkedlistvalues.Node
	expected []string
}

func newLinkedList(values ...string) *linkedlistvalues.Node {
	var (
		first *linkedlistvalues.Node
		last  *linkedlistvalues.Node
	)

	for _, value := range values {
		node := linkedlistvalues.NewNode(value)

		if first == nil {
			first = node
		}

		if last != nil {
			last.Next = node
		}

		last = node
	}

	return first
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList("a", "b", "c", "d"),
		expected: []string{"a", "b", "c", "d"},
	},
	{
		name:     "test_01",
		head:     newLinkedList("x", "y"),
		expected: []string{"x", "y"},
	},
	{
		name:     "test_02",
		head:     newLinkedList("q"),
		expected: []string{"q"},
	},
	{
		name:     "test_03",
		head:     newLinkedList(),
		expected: []string{},
	},
}
