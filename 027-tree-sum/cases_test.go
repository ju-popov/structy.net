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
			Val: 3,
			Left: &treesum.Node{
				Val: 11,
				Left: &treesum.Node{
					Val: 4,
				},
				Right: &treesum.Node{
					Val: -2,
				},
			},
			Right: &treesum.Node{
				Val: 4,
				Right: &treesum.Node{
					Val: 1,
				},
			},
		},
		expected: 21,
	},
	{
		name: "test_01",
		root: &treesum.Node{
			Val: 1,
			Left: &treesum.Node{
				Val: 6,
				Left: &treesum.Node{
					Val: 3,
				},
				Right: &treesum.Node{
					Val: -6,
					Left: &treesum.Node{
						Val: 2,
					},
				},
			},
			Right: &treesum.Node{
				Val: 0,
				Right: &treesum.Node{
					Val: 2,
					Right: &treesum.Node{
						Val: 2,
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
