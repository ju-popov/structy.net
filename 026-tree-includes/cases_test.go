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
			Val: "a",
			Left: &treeincludes.Node{
				Val: "b",
				Left: &treeincludes.Node{
					Val: "d",
				},
				Right: &treeincludes.Node{
					Val: "e",
				},
			},
			Right: &treeincludes.Node{
				Val: "c",
				Right: &treeincludes.Node{
					Val: "f",
				},
			},
		},
		target:   "e",
		expected: true,
	},
	{
		name: "test_01",
		root: &treeincludes.Node{
			Val: "a",
			Left: &treeincludes.Node{
				Val: "b",
				Left: &treeincludes.Node{
					Val: "d",
				},
				Right: &treeincludes.Node{
					Val: "e",
				},
			},
			Right: &treeincludes.Node{
				Val: "c",
				Right: &treeincludes.Node{
					Val: "f",
				},
			},
		},
		target:   "a",
		expected: true,
	},
	{
		name: "test_02",
		root: &treeincludes.Node{
			Val: "a",
			Left: &treeincludes.Node{
				Val: "b",
				Left: &treeincludes.Node{
					Val: "d",
				},
				Right: &treeincludes.Node{
					Val: "e",
				},
			},
			Right: &treeincludes.Node{
				Val: "c",
				Right: &treeincludes.Node{
					Val: "f",
				},
			},
		},
		target:   "n",
		expected: false,
	},
	{
		name: "test_03",
		root: &treeincludes.Node{
			Val: "a",
			Left: &treeincludes.Node{
				Val: "b",
				Left: &treeincludes.Node{
					Val: "d",
				},
				Right: &treeincludes.Node{
					Val: "e",
					Left: &treeincludes.Node{
						Val: "g",
					},
				},
			},
			Right: &treeincludes.Node{
				Val: "c",
				Right: &treeincludes.Node{
					Val: "f",
					Right: &treeincludes.Node{
						Val: "h",
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
			Val: "a",
			Left: &treeincludes.Node{
				Val: "b",
				Left: &treeincludes.Node{
					Val: "d",
				},
				Right: &treeincludes.Node{
					Val: "e",
					Left: &treeincludes.Node{
						Val: "g",
					},
				},
			},
			Right: &treeincludes.Node{
				Val: "c",
				Right: &treeincludes.Node{
					Val: "f",
					Right: &treeincludes.Node{
						Val: "h",
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
