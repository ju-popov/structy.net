package treepathfinder_test

import (
	treepathfinder "github.com/ju-popov/structy.net/030-tree-path-finder"
)

type testCase struct {
	name     string
	root     *treepathfinder.Node
	target   interface{}
	expected []interface{}
}

func rangeRightTree(max int) *treepathfinder.Node {
	var (
		root *treepathfinder.Node
		prev *treepathfinder.Node
	)

	for index := 0; index <= max; index++ {
		node := &treepathfinder.Node{
			Val: index,
		}

		if root == nil {
			root = node
		}

		if prev != nil {
			prev.Right = node
		}

		prev = node
	}

	return root
}

func rangeArray(max int) []interface{} {
	result := []interface{}{}

	for index := 0; index <= max; index++ {
		result = append(result, index)
	}

	return result
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &treepathfinder.Node{
			Val: "a",
			Left: &treepathfinder.Node{
				Val: "b",
				Left: &treepathfinder.Node{
					Val: "d",
				},
				Right: &treepathfinder.Node{
					Val: "e",
				},
			},
			Right: &treepathfinder.Node{
				Val: "c",
				Right: &treepathfinder.Node{
					Val: "f",
				},
			},
		},
		target:   "e",
		expected: []interface{}{"a", "b", "e"},
	},
	{
		name: "test_01",
		root: &treepathfinder.Node{
			Val: "a",
			Left: &treepathfinder.Node{
				Val: "b",
				Left: &treepathfinder.Node{
					Val: "d",
				},
				Right: &treepathfinder.Node{
					Val: "e",
				},
			},
			Right: &treepathfinder.Node{
				Val: "c",
				Right: &treepathfinder.Node{
					Val: "f",
				},
			},
		},
		target:   "p",
		expected: nil,
	},
	{
		name: "test_02",
		root: &treepathfinder.Node{
			Val: "a",
			Left: &treepathfinder.Node{
				Val: "b",
				Left: &treepathfinder.Node{
					Val: "d",
				},
				Right: &treepathfinder.Node{
					Val: "e",
					Left: &treepathfinder.Node{
						Val: "g",
					},
				},
			},
			Right: &treepathfinder.Node{
				Val: "c",
				Right: &treepathfinder.Node{
					Val: "f",
					Right: &treepathfinder.Node{
						Val: "h",
					},
				},
			},
		},
		target:   "c",
		expected: []interface{}{"a", "c"},
	},
	{
		name: "test_03",
		root: &treepathfinder.Node{
			Val: "a",
			Left: &treepathfinder.Node{
				Val: "b",
				Left: &treepathfinder.Node{
					Val: "d",
				},
				Right: &treepathfinder.Node{
					Val: "e",
					Left: &treepathfinder.Node{
						Val: "g",
					},
				},
			},
			Right: &treepathfinder.Node{
				Val: "c",
				Right: &treepathfinder.Node{
					Val: "f",
					Right: &treepathfinder.Node{
						Val: "h",
					},
				},
			},
		},
		target:   "h",
		expected: []interface{}{"a", "c", "f", "h"},
	},
	{
		name: "test_04",
		root: &treepathfinder.Node{
			Val: "x",
		},
		target:   "x",
		expected: []interface{}{"x"},
	},
	{
		name:     "test_05",
		root:     nil,
		target:   "x",
		expected: nil,
	},
	{
		name:     "test_06",
		root:     rangeRightTree(6000),
		target:   3451,
		expected: rangeArray(3451),
	},
}
