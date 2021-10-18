package treevaluecount_test

import (
	treevaluecount "github.com/ju-popov/structy.net/031-tree-value-count"
)

type testCase struct {
	name     string
	root     *treevaluecount.Node
	target   int64
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &treevaluecount.Node{
			Value: 12,
			Left: &treevaluecount.Node{
				Value: 6,
				Left: &treevaluecount.Node{
					Value: 4,
				},
				Right: &treevaluecount.Node{
					Value: 6,
				},
			},
			Right: &treevaluecount.Node{
				Value: 6,
				Right: &treevaluecount.Node{
					Value: 12,
				},
			},
		},
		target:   6,
		expected: 3,
	},
	{
		name: "test_01",
		root: &treevaluecount.Node{
			Value: 12,
			Left: &treevaluecount.Node{
				Value: 6,
				Left: &treevaluecount.Node{
					Value: 4,
				},
				Right: &treevaluecount.Node{
					Value: 6,
				},
			},
			Right: &treevaluecount.Node{
				Value: 6,
				Right: &treevaluecount.Node{
					Value: 12,
				},
			},
		},
		target:   12,
		expected: 2,
	},
	{
		name: "test_02",
		root: &treevaluecount.Node{
			Value: 7,
			Left: &treevaluecount.Node{
				Value: 5,
				Left: &treevaluecount.Node{
					Value: 1,
				},
				Right: &treevaluecount.Node{
					Value: 8,
					Left: &treevaluecount.Node{
						Value: 1,
					},
				},
			},
			Right: &treevaluecount.Node{
				Value: 1,
				Right: &treevaluecount.Node{
					Value: 7,
					Right: &treevaluecount.Node{
						Value: 1,
					},
				},
			},
		},
		target:   1,
		expected: 4,
	},
	{
		name: "test_03",
		root: &treevaluecount.Node{
			Value: 7,
			Left: &treevaluecount.Node{
				Value: 5,
				Left: &treevaluecount.Node{
					Value: 1,
				},
				Right: &treevaluecount.Node{
					Value: 8,
					Left: &treevaluecount.Node{
						Value: 1,
					},
				},
			},
			Right: &treevaluecount.Node{
				Value: 1,
				Right: &treevaluecount.Node{
					Value: 7,
					Right: &treevaluecount.Node{
						Value: 1,
					},
				},
			},
		},
		target:   9,
		expected: 0,
	},
	{
		name:     "test_04",
		root:     nil,
		target:   42,
		expected: 0,
	},
}
