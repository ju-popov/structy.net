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
			Value: "a",
			Left: &breadthfirstvalues.Node{
				Value: "b",
				Left: &breadthfirstvalues.Node{
					Value: "d",
				},
				Right: &breadthfirstvalues.Node{
					Value: "e",
				},
			},
			Right: &breadthfirstvalues.Node{
				Value: "c",
				Right: &breadthfirstvalues.Node{
					Value: "f",
				},
			},
		},
		expected: []string{"a", "b", "c", "d", "e", "f"},
	},
	{
		name: "test_01",
		root: &breadthfirstvalues.Node{
			Value: "a",
			Left: &breadthfirstvalues.Node{
				Value: "b",
				Left: &breadthfirstvalues.Node{
					Value: "d",
				},
				Right: &breadthfirstvalues.Node{
					Value: "e",
					Left: &breadthfirstvalues.Node{
						Value: "g",
					},
				},
			},
			Right: &breadthfirstvalues.Node{
				Value: "c",
				Right: &breadthfirstvalues.Node{
					Value: "f",
					Right: &breadthfirstvalues.Node{
						Value: "h",
					},
				},
			},
		},
		expected: []string{"a", "b", "c", "d", "e", "f", "g", "h"},
	},
	{
		name: "test_02",
		root: &breadthfirstvalues.Node{
			Value: "a",
		},
		expected: []string{"a"},
	},
	{
		name: "test_03",
		root: &breadthfirstvalues.Node{
			Value: "a",
			Right: &breadthfirstvalues.Node{
				Value: "b",
				Left: &breadthfirstvalues.Node{
					Value: "c",
					Left: &breadthfirstvalues.Node{
						Value: "x",
					},
					Right: &breadthfirstvalues.Node{
						Value: "d",
						Right: &breadthfirstvalues.Node{
							Value: "e",
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
