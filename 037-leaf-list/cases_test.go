package leaflist_test

import (
	leaflist "github.com/ju-popov/structy.net/037-leaf-list"
)

type testCase struct {
	name     string
	root     *leaflist.Node
	expected []interface{}
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &leaflist.Node{
			Val: "a",
			Left: &leaflist.Node{
				Val: "b",
				Left: &leaflist.Node{
					Val: "d",
				},
				Right: &leaflist.Node{
					Val: "e",
				},
			},
			Right: &leaflist.Node{
				Val: "c",
				Right: &leaflist.Node{
					Val: "f",
				},
			},
		},
		expected: []interface{}{"d", "e", "f"},
	},
	{
		name: "test_01",
		root: &leaflist.Node{
			Val: "a",
			Left: &leaflist.Node{
				Val: "b",
				Left: &leaflist.Node{
					Val: "d",
				},
				Right: &leaflist.Node{
					Val: "e",
					Left: &leaflist.Node{
						Val: "g",
					},
				},
			},
			Right: &leaflist.Node{
				Val: "c",
				Right: &leaflist.Node{
					Val: "f",
					Right: &leaflist.Node{
						Val: "h",
					},
				},
			},
		},
		expected: []interface{}{"d", "g", "h"},
	},
	{
		name: "test_02",
		root: &leaflist.Node{
			Val: 5,
			Left: &leaflist.Node{
				Val: 11,
				Left: &leaflist.Node{
					Val: 20,
				},
				Right: &leaflist.Node{
					Val: 15,
					Left: &leaflist.Node{
						Val: 1,
					},
					Right: &leaflist.Node{
						Val: 3,
					},
				},
			},
			Right: &leaflist.Node{
				Val: 54,
			},
		},
		expected: []interface{}{20, 1, 3, 54},
	},
	{
		name: "test_03",
		root: &leaflist.Node{
			Val: "x",
		},
		expected: []interface{}{"x"},
	},
	{
		name:     "test_04",
		root:     nil,
		expected: []interface{}{},
	},
}
