package breadthfirstvalues_test

import (
	breadthfirstvalues "github.com/ju-popov/structy.net/025-breadth-first-values"
)

type testCase struct {
	name     string
	root     *breadthfirstvalues.Node
	expected []string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &breadthfirstvalues.Node{
			Val: "a",
			Left: &breadthfirstvalues.Node{
				Val: "b",
				Left: &breadthfirstvalues.Node{
					Val: "d",
				},
				Right: &breadthfirstvalues.Node{
					Val: "e",
				},
			},
			Right: &breadthfirstvalues.Node{
				Val: "c",
				Right: &breadthfirstvalues.Node{
					Val: "f",
				},
			},
		},
		expected: []string{"a", "b", "c", "d", "e", "f"},
	},
	{
		name: "test_01",
		root: &breadthfirstvalues.Node{
			Val: "a",
			Left: &breadthfirstvalues.Node{
				Val: "b",
				Left: &breadthfirstvalues.Node{
					Val: "d",
				},
				Right: &breadthfirstvalues.Node{
					Val: "e",
					Left: &breadthfirstvalues.Node{
						Val: "g",
					},
				},
			},
			Right: &breadthfirstvalues.Node{
				Val: "c",
				Right: &breadthfirstvalues.Node{
					Val: "f",
					Right: &breadthfirstvalues.Node{
						Val: "h",
					},
				},
			},
		},
		expected: []string{"a", "b", "c", "d", "e", "f", "g", "h"},
	},
	{
		name: "test_02",
		root: &breadthfirstvalues.Node{
			Val: "a",
		},
		expected: []string{"a"},
	},
	{
		name: "test_03",
		root: &breadthfirstvalues.Node{
			Val: "a",
			Right: &breadthfirstvalues.Node{
				Val: "b",
				Left: &breadthfirstvalues.Node{
					Val: "c",
					Left: &breadthfirstvalues.Node{
						Val: "x",
					},
					Right: &breadthfirstvalues.Node{
						Val: "d",
						Right: &breadthfirstvalues.Node{
							Val: "e",
						},
					},
				},
			},
		},
		expected: []string{"a", "b", "c", "x", "d", "e"},
	},
	{
		name:     "test_04",
		root:     nil,
		expected: []string{},
	},
}
