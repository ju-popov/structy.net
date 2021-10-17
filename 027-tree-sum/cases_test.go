package treesum_test

import (
	treesum "github.com/ju-popov/structy.net/027-tree-sum"
)

type testCase struct {
	name     string
	root     *treesum.Node
	expected int64
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &treesum.Node{
			Value: 3,
			Left: &treesum.Node{
				Value: 11,
				Left: &treesum.Node{
					Value: 4,
				},
				Right: &treesum.Node{
					Value: -2,
				},
			},
			Right: &treesum.Node{
				Value: 4,
				Right: &treesum.Node{
					Value: 1,
				},
			},
		},
		expected: 21,
	},
	{
		name: "test_01",
		root: &treesum.Node{
			Value: 1,
			Left: &treesum.Node{
				Value: 6,
				Left: &treesum.Node{
					Value: 3,
				},
				Right: &treesum.Node{
					Value: -6,
					Left: &treesum.Node{
						Value: 2,
					},
				},
			},
			Right: &treesum.Node{
				Value: 0,
				Right: &treesum.Node{
					Value: 2,
					Right: &treesum.Node{
						Value: 2,
					},
				},
			},
		},
		expected: 10,
	},
	{
		name:     "test_02",
		root:     nil,
		expected: 0,
	},
}
