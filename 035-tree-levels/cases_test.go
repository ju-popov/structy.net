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
			Value: "a",
			Left: &treelevels.Node{
				Value: "b",
				Left: &treelevels.Node{
					Value: "d",
				},
				Right: &treelevels.Node{
					Value: "e",
				},
			},
			Right: &treelevels.Node{
				Value: "c",
				Right: &treelevels.Node{
					Value: "f",
				},
			},
		},
		expected: [][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}},
	},
	{
		name: "test_01",
		root: &treelevels.Node{
			Value: "a",
			Left: &treelevels.Node{
				Value: "b",
				Left: &treelevels.Node{
					Value: "d",
				},
				Right: &treelevels.Node{
					Value: "e",
					Left: &treelevels.Node{
						Value: "g",
					},
					Right: &treelevels.Node{
						Value: "h",
					},
				},
			},
			Right: &treelevels.Node{
				Value: "c",
				Right: &treelevels.Node{
					Value: "f",
					Left: &treelevels.Node{
						Value: "i",
					},
				},
			},
		},
		expected: [][]string{{"a"}, {"b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}},
	},
	{
		name: "test_02",
		root: &treelevels.Node{
			Value: "q",
			Left: &treelevels.Node{
				Value: "r",
				Right: &treelevels.Node{
					Value: "t",
					Left: &treelevels.Node{
						Value: "u",
						Left: &treelevels.Node{
							Value: "v",
						},
					},
				},
			},
			Right: &treelevels.Node{
				Value: "s",
			},
		},
		expected: [][]string{{"q"}, {"r", "s"}, {"t"}, {"u"}, {"v"}},
	},
	{
		name: "test_03",
		root: nil,
		expected: [][]string{},
	},
}
