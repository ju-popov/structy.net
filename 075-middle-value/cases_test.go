package middlevalue_test

import middlevalue "github.com/ju-popov/structy.net/075-middle-value"

type testCase struct {
	name     string
	head     *middlevalue.Node
	expected string
}

func newLinkedList(values ...string) *middlevalue.Node {
	dummyHead := middlevalue.NewNode("")
	current := dummyHead

	for _, value := range values {
		current.Next = middlevalue.NewNode(value)
		current = current.Next
	}

	return dummyHead.Next
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		head:     newLinkedList("a", "b", "c", "d", "e"),
		expected: "c",
	},
	{
		name:     "test_01",
		head:     newLinkedList("a", "b", "c", "d", "e", "f"),
		expected: "d",
	},
	{
		name:     "test_02",
		head:     newLinkedList("x", "y", "z"),
		expected: "y",
	},
	{
		name:     "test_03",
		head:     newLinkedList("x", "y"),
		expected: "y",
	},
	{
		name:     "test_04",
		head:     newLinkedList("q"),
		expected: "q",
	},
}
