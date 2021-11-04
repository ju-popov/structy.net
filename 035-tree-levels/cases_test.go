package treelevels_test

import (
	treelevels "github.com/ju-popov/structy.net/035-tree-levels"
)

type testCase struct {
	name     string
	root     *treelevels.Node
	expected [][]string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &treelevels.Node{
			Val: "a",
			Left: &treelevels.Node{
				Val: "b",
				Left: &treelevels.Node{
					Val: "d",
				},
				Right: &treelevels.Node{
					Val: "e",
				},
			},
			Right: &treelevels.Node{
				Val: "c",
				Right: &treelevels.Node{
					Val: "f",
				},
			},
		},
		expected: [][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}},
	},
	{
		name: "test_01",
		root: &treelevels.Node{
			Val: "a",
			Left: &treelevels.Node{
				Val: "b",
				Left: &treelevels.Node{
					Val: "d",
				},
				Right: &treelevels.Node{
					Val: "e",
					Left: &treelevels.Node{
						Val: "g",
					},
					Right: &treelevels.Node{
						Val: "h",
					},
				},
			},
			Right: &treelevels.Node{
				Val: "c",
				Right: &treelevels.Node{
					Val: "f",
					Left: &treelevels.Node{
						Val: "i",
					},
				},
			},
		},
		expected: [][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}},
	},
	{
		name: "test_02",
		root: &treelevels.Node{
			Val: "q",
			Left: &treelevels.Node{
				Val: "r",
				Right: &treelevels.Node{
					Val: "t",
					Left: &treelevels.Node{
						Val: "u",
						Left: &treelevels.Node{
							Val: "v",
						},
					},
				},
			},
			Right: &treelevels.Node{
				Val: "s",
			},
		},
		expected: [][]string{{"q"}, {"r", "s"}, {"t"}, {"u"}, {"v"}},
	},
	{
		name:     "test_03",
		root:     nil,
		expected: [][]string{},
	},
}
