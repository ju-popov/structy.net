package alltreepaths_test

import (
	alltreepaths "github.com/ju-popov/structy.net/034-all-tree-paths"
)

type testCase struct {
	name     string
	root     *alltreepaths.Node
	expected [][]string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &alltreepaths.Node{
			Val: "a",
			Left: &alltreepaths.Node{
				Val: "b",
				Left: &alltreepaths.Node{
					Val: "d",
				},
				Right: &alltreepaths.Node{
					Val: "e",
				},
			},
			Right: &alltreepaths.Node{
				Val: "c",
				Right: &alltreepaths.Node{
					Val: "f",
				},
			},
		},
		expected: [][]string{{"a", "b", "d"}, {"a", "b", "e"}, {"a", "c", "f"}},
	},
	{
		name: "test_01",
		root: &alltreepaths.Node{
			Val: "a",
			Left: &alltreepaths.Node{
				Val: "b",
				Left: &alltreepaths.Node{
					Val: "d",
				},
				Right: &alltreepaths.Node{
					Val: "e",
					Left: &alltreepaths.Node{
						Val: "g",
					},
					Right: &alltreepaths.Node{
						Val: "h",
					},
				},
			},
			Right: &alltreepaths.Node{
				Val: "c",
				Right: &alltreepaths.Node{
					Val: "f",
					Left: &alltreepaths.Node{
						Val: "i",
					},
				},
			},
		},
		expected: [][]string{{"a", "b", "d"}, {"a", "b", "e", "g"}, {"a", "b", "e", "h"}, {"a", "c", "f", "i"}},
	},
	{
		name: "test_02",
		root: &alltreepaths.Node{
			Val: "q",
			Left: &alltreepaths.Node{
				Val: "r",
				Right: &alltreepaths.Node{
					Val: "t",
					Left: &alltreepaths.Node{
						Val: "u",
						Left: &alltreepaths.Node{
							Val: "v",
						},
					},
				},
			},
			Right: &alltreepaths.Node{
				Val: "s",
			},
		},
		expected: [][]string{{"q", "r", "t", "u", "v"}, {"q", "s"}},
	},
	{
		name: "test_03",
		root: &alltreepaths.Node{
			Val: "z",
		},
		expected: [][]string{{"z"}},
	},
}
