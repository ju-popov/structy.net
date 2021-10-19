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
			Value: "a",
			Left: &howhigh.Node{
				Value: "b",
				Left: &howhigh.Node{
					Value: "d",
				},
				Right: &howhigh.Node{
					Value: "e",
				},
			},
			Right: &howhigh.Node{
				Value: "c",
				Right: &howhigh.Node{
					Value: "f",
				},
			},
		},
		expected: 2,
	},
	{
		name: "test_01",
		root: &howhigh.Node{
			Value: "a",
			Left: &howhigh.Node{
				Value: "b",
				Left: &howhigh.Node{
					Value: "d",
				},
				Right: &howhigh.Node{
					Value: "e",
					Left: &howhigh.Node{
						Value: "g",
					},
				},
			},
			Right: &howhigh.Node{
				Value: "c",
				Right: &howhigh.Node{
					Value: "f",
				},
			},
		},
		expected: 3,
	},
	{
		name: "test_02",
		root: &howhigh.Node{
			Value: "a",
			Right: &howhigh.Node{
				Value: "c",
			},
		},
		expected: 1,
	},
	{
		name: "test_03",
		root: &howhigh.Node{
			Value: "a",
		},
		expected: 0,
	},
	{
		name:     "test_04",
		root:     nil,
		expected: -1,
	},
}
