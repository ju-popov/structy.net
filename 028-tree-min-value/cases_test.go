package treeminvalue_test

import (
	treeminvalue "github.com/ju-popov/structy.net/028-tree-min-value"
)

type testCase struct {
	name     string
	root     *treeminvalue.Node
	expected int64
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &treeminvalue.Node{
			Value: 3,
			Left: &treeminvalue.Node{
				Value: 11,
				Left: &treeminvalue.Node{
					Value: 4,
				},
				Right: &treeminvalue.Node{
					Value: -2,
				},
			},
			Right: &treeminvalue.Node{
				Value: 4,
				Right: &treeminvalue.Node{
					Value: 1,
				},
			},
		},
		expected: -2,
	},
	{
		name: "test_01",
		root: &treeminvalue.Node{
			Value: 5,
			Left: &treeminvalue.Node{
				Value: 11,
				Left: &treeminvalue.Node{
					Value: 4,
				},
				Right: &treeminvalue.Node{
					Value: 15,
				},
			},
			Right: &treeminvalue.Node{
				Value: 3,
				Right: &treeminvalue.Node{
					Value: 12,
				},
			},
		},
		expected: 3,
	},
	{
		name: "test_02",
		root: &treeminvalue.Node{
			Value: -1,
			Left: &treeminvalue.Node{
				Value: -6,
				Left: &treeminvalue.Node{
					Value: -3,
				},
				Right: &treeminvalue.Node{
					Value: -4,
					Left: &treeminvalue.Node{
						Value: -2,
					},
				},
			},
			Right: &treeminvalue.Node{
				Value: -5,
				Right: &treeminvalue.Node{
					Value: -13,
					Right: &treeminvalue.Node{
						Value: -2,
					},
				},
			},
		},
		expected: -13,
	},
	{
		name: "test_03",
		root: &treeminvalue.Node{
			Value: 42,
		},
		expected: 42,
	},
}
