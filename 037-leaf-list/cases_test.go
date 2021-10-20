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
			Value: "a",
			Left: &leaflist.Node{
				Value: "b",
				Left: &leaflist.Node{
					Value: "d",
				},
				Right: &leaflist.Node{
					Value: "e",
				},
			},
			Right: &leaflist.Node{
				Value: "c",
				Right: &leaflist.Node{
					Value: "f",
				},
			},
		},
		expected: []interface{}{"d", "e", "f"},
	},
	{
		name: "test_01",
		root: &leaflist.Node{
			Value: "a",
			Left: &leaflist.Node{
				Value: "b",
				Left: &leaflist.Node{
					Value: "d",
				},
				Right: &leaflist.Node{
					Value: "e",
					Left: &leaflist.Node{
						Value: "g",
					},
				},
			},
			Right: &leaflist.Node{
				Value: "c",
				Right: &leaflist.Node{
					Value: "f",
					Right: &leaflist.Node{
						Value: "h",
					},
				},
			},
		},
		expected: []interface{}{"d", "g", "h"},
	},
	{
		name: "test_02",
		root: &leaflist.Node{
			Value: 5,
			Left: &leaflist.Node{
				Value: 11,
				Left: &leaflist.Node{
					Value: 20,
				},
				Right: &leaflist.Node{
					Value: 15,
					Left: &leaflist.Node{
						Value: 1,
					},
					Right: &leaflist.Node{
						Value: 3,
					},
				},
			},
			Right: &leaflist.Node{
				Value: 54,
			},
		},
		expected: []interface{}{20, 1, 3, 54},
	},
	{
		name: "test_03",
		root: &leaflist.Node{
			Value: "x",
		},
		expected: []interface{}{"x"},
	},
	{
		name:     "test_04",
		root:     nil,
		expected: []interface{}{},
	},
}
