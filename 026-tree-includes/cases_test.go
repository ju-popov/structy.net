package treeincludes_test

import (
	treeincludes "github.com/ju-popov/structy.net/026-tree-includes"
)

type testCase struct {
	name     string
	root     *treeincludes.Node
	target   string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &treeincludes.Node{
			Value: "a",
			Left: &treeincludes.Node{
				Value: "b",
				Left: &treeincludes.Node{
					Value: "d",
				},
				Right: &treeincludes.Node{
					Value: "e",
				},
			},
			Right: &treeincludes.Node{
				Value: "c",
				Right: &treeincludes.Node{
					Value: "f",
				},
			},
		},
		target:   "e",
		expected: true,
	},
	{
		name: "test_01",
		root: &treeincludes.Node{
			Value: "a",
			Left: &treeincludes.Node{
				Value: "b",
				Left: &treeincludes.Node{
					Value: "d",
				},
				Right: &treeincludes.Node{
					Value: "e",
				},
			},
			Right: &treeincludes.Node{
				Value: "c",
				Right: &treeincludes.Node{
					Value: "f",
				},
			},
		},
		target:   "a",
		expected: true,
	},
	{
		name: "test_02",
		root: &treeincludes.Node{
			Value: "a",
			Left: &treeincludes.Node{
				Value: "b",
				Left: &treeincludes.Node{
					Value: "d",
				},
				Right: &treeincludes.Node{
					Value: "e",
				},
			},
			Right: &treeincludes.Node{
				Value: "c",
				Right: &treeincludes.Node{
					Value: "f",
				},
			},
		},
		target:   "n",
		expected: false,
	},
	{
		name: "test_03",
		root: &treeincludes.Node{
			Value: "a",
			Left: &treeincludes.Node{
				Value: "b",
				Left: &treeincludes.Node{
					Value: "d",
				},
				Right: &treeincludes.Node{
					Value: "e",
					Left: &treeincludes.Node{
						Value: "g",
					},
				},
			},
			Right: &treeincludes.Node{
				Value: "c",
				Right: &treeincludes.Node{
					Value: "f",
					Right: &treeincludes.Node{
						Value: "h",
					},
				},
			},
		},
		target:   "f",
		expected: true,
	},
	{
		name: "test_04",
		root: &treeincludes.Node{
			Value: "a",
			Left: &treeincludes.Node{
				Value: "b",
				Left: &treeincludes.Node{
					Value: "d",
				},
				Right: &treeincludes.Node{
					Value: "e",
					Left: &treeincludes.Node{
						Value: "g",
					},
				},
			},
			Right: &treeincludes.Node{
				Value: "c",
				Right: &treeincludes.Node{
					Value: "f",
					Right: &treeincludes.Node{
						Value: "h",
					},
				},
			},
		},
		target:   "p",
		expected: false,
	},
	{
		name:     "test_05",
		root:     nil,
		target:   "b",
		expected: false,
	},
}
