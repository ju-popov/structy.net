package howhigh_test

import (
	howhigh "github.com/ju-popov/structy.net/032-how-high"
)

type testCase struct {
	name     string
	root     *howhigh.Node
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &howhigh.Node{
			Val: "a",
			Left: &howhigh.Node{
				Val: "b",
				Left: &howhigh.Node{
					Val: "d",
				},
				Right: &howhigh.Node{
					Val: "e",
				},
			},
			Right: &howhigh.Node{
				Val: "c",
				Right: &howhigh.Node{
					Val: "f",
				},
			},
		},
		expected: 2,
	},
	{
		name: "test_01",
		root: &howhigh.Node{
			Val: "a",
			Left: &howhigh.Node{
				Val: "b",
				Left: &howhigh.Node{
					Val: "d",
				},
				Right: &howhigh.Node{
					Val: "e",
					Left: &howhigh.Node{
						Val: "g",
					},
				},
			},
			Right: &howhigh.Node{
				Val: "c",
				Right: &howhigh.Node{
					Val: "f",
				},
			},
		},
		expected: 3,
	},
	{
		name: "test_02",
		root: &howhigh.Node{
			Val: "a",
			Right: &howhigh.Node{
				Val: "c",
			},
		},
		expected: 1,
	},
	{
		name: "test_03",
		root: &howhigh.Node{
			Val: "a",
		},
		expected: 0,
	},
	{
		name:     "test_04",
		root:     nil,
		expected: -1,
	},
}
