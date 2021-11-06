package leftynodes_test

import leftynodes "github.com/ju-popov/structy.net/079-lefty-nodes"

type testCase struct {
	name     string
	root     *leftynodes.Node
	expected []string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &leftynodes.Node{
			Val: "a",
			Left: &leftynodes.Node{
				Val: "b",
				Left: &leftynodes.Node{
					Val: "d",
				},
				Right: &leftynodes.Node{
					Val: "e",
					Left: &leftynodes.Node{
						Val: "g",
					},
					Right: &leftynodes.Node{
						Val: "h",
					},
				},
			},
			Right: &leftynodes.Node{
				Val: "c",
				Right: &leftynodes.Node{
					Val: "f",
				},
			},
		},
		expected: []string{"a", "b", "d", "g"},
	},
	{
		name: "test_01",
		root: &leftynodes.Node{
			Val: "u",
			Left: &leftynodes.Node{
				Val: "t",
			},
			Right: &leftynodes.Node{
				Val: "s",
				Right: &leftynodes.Node{
					Val: "r",
					Left: &leftynodes.Node{
						Val: "q",
					},
					Right: &leftynodes.Node{
						Val: "p",
					},
				},
			},
		},
		expected: []string{"u", "t", "r", "q"},
	},
	{
		name: "test_02",
		root: &leftynodes.Node{
			Val: "l",
			Left: &leftynodes.Node{
				Val: "m",
			},
			Right: &leftynodes.Node{
				Val: "n",
				Left: &leftynodes.Node{
					Val: "o",
					Left: &leftynodes.Node{
						Val: "q",
					},
					Right: &leftynodes.Node{
						Val: "r",
					},
				},
				Right: &leftynodes.Node{
					Val: "p",
					Left: &leftynodes.Node{
						Val: "s",
					},
					Right: &leftynodes.Node{
						Val: "t",
					},
				},
			},
		},
		expected: []string{"l", "m", "o", "q"},
	},
	{
		name: "test_03",
		root: &leftynodes.Node{
			Val: "n",
			Left: &leftynodes.Node{
				Val: "y",
			},
			Right: &leftynodes.Node{
				Val: "c",
			},
		},
		expected: []string{"n", "y"},
	},
	{
		name: "test_04",
		root: &leftynodes.Node{
			Val: "i",
			Right: &leftynodes.Node{
				Val: "n",
				Left: &leftynodes.Node{
					Val: "s",
				},
				Right: &leftynodes.Node{
					Val: "t",
				},
			},
		},
		expected: []string{"i", "n", "s"},
	},
	{
		name:     "test_05",
		root:     nil,
		expected: []string{},
	},
}
