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
			Value: "a",
			Left: &alltreepaths.Node{
				Value: "b",
				Left: &alltreepaths.Node{
					Value: "d",
				},
				Right: &alltreepaths.Node{
					Value: "e",
				},
			},
			Right: &alltreepaths.Node{
				Value: "c",
				Right: &alltreepaths.Node{
					Value: "f",
				},
			},
		},
		expected: [][]string{{"a", "b", "d"}, {"a", "b", "e"}, {"a", "c", "f"}},
	},
	{
		name: "test_01",
		root: &alltreepaths.Node{
			Value: "a",
			Left: &alltreepaths.Node{
				Value: "b",
				Left: &alltreepaths.Node{
					Value: "d",
				},
				Right: &alltreepaths.Node{
					Value: "e",
					Left: &alltreepaths.Node{
						Value: "g",
					},
					Right: &alltreepaths.Node{
						Value: "h",
					},
				},
			},
			Right: &alltreepaths.Node{
				Value: "c",
				Right: &alltreepaths.Node{
					Value: "f",
					Left: &alltreepaths.Node{
						Value: "i",
					},
				},
			},
		},
		expected: [][]string{{"a", "b", "d"}, {"a", "b", "e", "g"}, {"a", "b", "e", "h"}, {"a", "c", "f", "i"}},
	},
	{
		name: "test_02",
		root: &alltreepaths.Node{
			Value: "q",
			Left: &alltreepaths.Node{
				Value: "r",
				Right: &alltreepaths.Node{
					Value: "t",
					Left: &alltreepaths.Node{
						Value: "u",
						Left: &alltreepaths.Node{
							Value: "v",
						},
					},
				},
			},
			Right: &alltreepaths.Node{
				Value: "s",
			},
		},
		expected: [][]string{{"q", "r", "t", "u", "v"}, {"q", "s"}},
	},
	{
		name: "test_03",
		root: &alltreepaths.Node{
			Value: "z",
		},
		expected: [][]string{{"z"}},
	},
}
