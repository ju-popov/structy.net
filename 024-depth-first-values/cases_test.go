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
			Value: "a",
			Left: &depthfirstvalues.Node{
				Value: "b",
				Left: &depthfirstvalues.Node{
					Value: "d",
				},
				Right: &depthfirstvalues.Node{
					Value: "e",
				},
			},
			Right: &depthfirstvalues.Node{
				Value: "c",
				Right: &depthfirstvalues.Node{
					Value: "f",
				},
			},
		},
		expected: []string{"a", "b", "d", "e", "c", "f"},
	},
	{
		name: "test_01",
		root: &depthfirstvalues.Node{
			Value: "a",
			Left: &depthfirstvalues.Node{
				Value: "b",
				Left: &depthfirstvalues.Node{
					Value: "d",
				},
				Right: &depthfirstvalues.Node{
					Value: "e",
					Left: &depthfirstvalues.Node{
						Value: "g",
					},
				},
			},
			Right: &depthfirstvalues.Node{
				Value: "c",
				Right: &depthfirstvalues.Node{
					Value: "f",
				},
			},
		},
		expected: []string{"a", "b", "d", "e", "g", "c", "f"},
	},
	{
		name: "test_02",
		root: &depthfirstvalues.Node{
			Value: "a",
		},
		expected: []string{"a"},
	},
	{
		name: "test_03",
		root: &depthfirstvalues.Node{
			Value: "a",
			Right: &depthfirstvalues.Node{
				Value: "b",
				Left: &depthfirstvalues.Node{
					Value: "c",
					Right: &depthfirstvalues.Node{
						Value: "d",
						Right: &depthfirstvalues.Node{
							Value: "e",
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
