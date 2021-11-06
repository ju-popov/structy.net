package fliptree_test

import fliptree "github.com/ju-popov/structy.net/078-flip-tree"

type testCase struct {
	name     string
	root     *fliptree.Node
	expected *fliptree.Node
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &fliptree.Node{
			Val: "a",
			Left: &fliptree.Node{
				Val: "b",
				Left: &fliptree.Node{
					Val: "d",
				},
				Right: &fliptree.Node{
					Val: "e",
					Left: &fliptree.Node{
						Val: "g",
					},
					Right: &fliptree.Node{
						Val: "h",
					},
				},
			},
			Right: &fliptree.Node{
				Val: "c",
				Right: &fliptree.Node{
					Val: "f",
				},
			},
		},
		expected: &fliptree.Node{
			Val: "a",
			Left: &fliptree.Node{
				Val: "c",
				Left: &fliptree.Node{
					Val: "f",
				},
			},
			Right: &fliptree.Node{
				Val: "b",
				Left: &fliptree.Node{
					Val: "e",
					Left: &fliptree.Node{
						Val: "h",
					},
					Right: &fliptree.Node{
						Val: "g",
					},
				},
				Right: &fliptree.Node{
					Val: "d",
				},
			},
		},
	},
	{
		name: "test_01",
		root: &fliptree.Node{
			Val: "u",
			Left: &fliptree.Node{
				Val: "t",
			},
			Right: &fliptree.Node{
				Val: "s",
				Right: &fliptree.Node{
					Val: "r",
					Left: &fliptree.Node{
						Val: "q",
					},
					Right: &fliptree.Node{
						Val: "p",
					},
				},
			},
		},
		expected: &fliptree.Node{
			Val: "u",
			Left: &fliptree.Node{
				Val: "s",
				Left: &fliptree.Node{
					Val: "r",
					Left: &fliptree.Node{
						Val: "p",
					},
					Right: &fliptree.Node{
						Val: "q",
					},
				},
			},
			Right: &fliptree.Node{
				Val: "t",
			},
		},
	},
	{
		name: "test_02",
		root: &fliptree.Node{
			Val: "l",
			Left: &fliptree.Node{
				Val: "m",
			},
			Right: &fliptree.Node{
				Val: "n",
				Left: &fliptree.Node{
					Val: "o",
					Left: &fliptree.Node{
						Val: "q",
					},
					Right: &fliptree.Node{
						Val: "r",
					},
				},
				Right: &fliptree.Node{
					Val: "p",
					Left: &fliptree.Node{
						Val: "s",
					},
					Right: &fliptree.Node{
						Val: "t",
					},
				},
			},
		},
		expected: &fliptree.Node{
			Val: "l",
			Left: &fliptree.Node{
				Val: "n",
				Left: &fliptree.Node{
					Val: "p",
					Left: &fliptree.Node{
						Val: "t",
					},
					Right: &fliptree.Node{
						Val: "s",
					},
				},
				Right: &fliptree.Node{
					Val: "o",
					Left: &fliptree.Node{
						Val: "r",
					},
					Right: &fliptree.Node{
						Val: "q",
					},
				},
			},
			Right: &fliptree.Node{
				Val: "m",
			},
		},
	},
	{
		name: "test_03",
		root: &fliptree.Node{
			Val: "n",
			Left: &fliptree.Node{
				Val: "y",
			},
			Right: &fliptree.Node{
				Val: "c",
			},
		},
		expected: &fliptree.Node{
			Val: "n",
			Left: &fliptree.Node{
				Val: "c",
			},
			Right: &fliptree.Node{
				Val: "y",
			},
		},
	},
}
