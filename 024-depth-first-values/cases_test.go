package depthfirstvalues_test

import (
	depthfirstvalues "github.com/ju-popov/structy.net/024-depth-first-values"
)

type testCase struct {
	name     string
	root     *depthfirstvalues.Node
	expected []string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &depthfirstvalues.Node{
			Val: "a",
			Left: &depthfirstvalues.Node{
				Val: "b",
				Left: &depthfirstvalues.Node{
					Val: "d",
				},
				Right: &depthfirstvalues.Node{
					Val: "e",
				},
			},
			Right: &depthfirstvalues.Node{
				Val: "c",
				Right: &depthfirstvalues.Node{
					Val: "f",
				},
			},
		},
		expected: []string{"a", "b", "d", "e", "c", "f"},
	},
	{
		name: "test_01",
		root: &depthfirstvalues.Node{
			Val: "a",
			Left: &depthfirstvalues.Node{
				Val: "b",
				Left: &depthfirstvalues.Node{
					Val: "d",
				},
				Right: &depthfirstvalues.Node{
					Val: "e",
					Left: &depthfirstvalues.Node{
						Val: "g",
					},
				},
			},
			Right: &depthfirstvalues.Node{
				Val: "c",
				Right: &depthfirstvalues.Node{
					Val: "f",
				},
			},
		},
		expected: []string{"a", "b", "d", "e", "g", "c", "f"},
	},
	{
		name: "test_02",
		root: &depthfirstvalues.Node{
			Val: "a",
		},
		expected: []string{"a"},
	},
	{
		name: "test_03",
		root: &depthfirstvalues.Node{
			Val: "a",
			Right: &depthfirstvalues.Node{
				Val: "b",
				Left: &depthfirstvalues.Node{
					Val: "c",
					Right: &depthfirstvalues.Node{
						Val: "d",
						Right: &depthfirstvalues.Node{
							Val: "e",
						},
					},
				},
			},
		},
		expected: []string{"a", "b", "c", "d", "e"},
	},
	{
		name:     "test_04",
		root:     nil,
		expected: []string{},
	},
}
