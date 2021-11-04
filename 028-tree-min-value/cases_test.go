package treeminvalue_test

import (
	treeminvalue "github.com/ju-popov/structy.net/028-tree-min-value"
)

type testCase struct {
	name     string
	root     *treeminvalue.Node
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &treeminvalue.Node{
			Val: 3,
			Left: &treeminvalue.Node{
				Val: 11,
				Left: &treeminvalue.Node{
					Val: 4,
				},
				Right: &treeminvalue.Node{
					Val: -2,
				},
			},
			Right: &treeminvalue.Node{
				Val: 4,
				Right: &treeminvalue.Node{
					Val: 1,
				},
			},
		},
		expected: -2,
	},
	{
		name: "test_01",
		root: &treeminvalue.Node{
			Val: 5,
			Left: &treeminvalue.Node{
				Val: 11,
				Left: &treeminvalue.Node{
					Val: 4,
				},
				Right: &treeminvalue.Node{
					Val: 15,
				},
			},
			Right: &treeminvalue.Node{
				Val: 3,
				Right: &treeminvalue.Node{
					Val: 12,
				},
			},
		},
		expected: 3,
	},
	{
		name: "test_02",
		root: &treeminvalue.Node{
			Val: -1,
			Left: &treeminvalue.Node{
				Val: -6,
				Left: &treeminvalue.Node{
					Val: -3,
				},
				Right: &treeminvalue.Node{
					Val: -4,
					Left: &treeminvalue.Node{
						Val: -2,
					},
				},
			},
			Right: &treeminvalue.Node{
				Val: -5,
				Right: &treeminvalue.Node{
					Val: -13,
					Right: &treeminvalue.Node{
						Val: -2,
					},
				},
			},
		},
		expected: -13,
	},
	{
		name: "test_03",
		root: &treeminvalue.Node{
			Val: 42,
		},
		expected: 42,
	},
}
