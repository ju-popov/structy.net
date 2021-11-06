package lowestcommonancestor_test

import lowestcommonancestor "github.com/ju-popov/structy.net/077-lowest-common-ancestor"

type testCase struct {
	name     string
	root     *lowestcommonancestor.Node
	val1     string
	val2     string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &lowestcommonancestor.Node{
			Val: "a",
			Left: &lowestcommonancestor.Node{
				Val: "b",
				Left: &lowestcommonancestor.Node{
					Val: "d",
				},
				Right: &lowestcommonancestor.Node{
					Val: "e",
					Left: &lowestcommonancestor.Node{
						Val: "g",
					},
					Right: &lowestcommonancestor.Node{
						Val: "h",
					},
				},
			},
			Right: &lowestcommonancestor.Node{
				Val: "c",
				Right: &lowestcommonancestor.Node{
					Val: "f",
				},
			},
		},
		val1:     "d",
		val2:     "h",
		expected: "b",
	},
	{
		name: "test_01",
		root: &lowestcommonancestor.Node{
			Val: "a",
			Left: &lowestcommonancestor.Node{
				Val: "b",
				Left: &lowestcommonancestor.Node{
					Val: "d",
				},
				Right: &lowestcommonancestor.Node{
					Val: "e",
					Left: &lowestcommonancestor.Node{
						Val: "g",
					},
					Right: &lowestcommonancestor.Node{
						Val: "h",
					},
				},
			},
			Right: &lowestcommonancestor.Node{
				Val: "c",
				Right: &lowestcommonancestor.Node{
					Val: "f",
				},
			},
		},
		val1:     "d",
		val2:     "g",
		expected: "b",
	},
	{
		name: "test_02",
		root: &lowestcommonancestor.Node{
			Val: "a",
			Left: &lowestcommonancestor.Node{
				Val: "b",
				Left: &lowestcommonancestor.Node{
					Val: "d",
				},
				Right: &lowestcommonancestor.Node{
					Val: "e",
					Left: &lowestcommonancestor.Node{
						Val: "g",
					},
					Right: &lowestcommonancestor.Node{
						Val: "h",
					},
				},
			},
			Right: &lowestcommonancestor.Node{
				Val: "c",
				Right: &lowestcommonancestor.Node{
					Val: "f",
				},
			},
		},
		val1:     "g",
		val2:     "c",
		expected: "a",
	},
	{
		name: "test_03",
		root: &lowestcommonancestor.Node{
			Val: "a",
			Left: &lowestcommonancestor.Node{
				Val: "b",
				Left: &lowestcommonancestor.Node{
					Val: "d",
				},
				Right: &lowestcommonancestor.Node{
					Val: "e",
					Left: &lowestcommonancestor.Node{
						Val: "g",
					},
					Right: &lowestcommonancestor.Node{
						Val: "h",
					},
				},
			},
			Right: &lowestcommonancestor.Node{
				Val: "c",
				Right: &lowestcommonancestor.Node{
					Val: "f",
				},
			},
		},
		val1:     "b",
		val2:     "g",
		expected: "b",
	},
	{
		name: "test_04",
		root: &lowestcommonancestor.Node{
			Val: "a",
			Left: &lowestcommonancestor.Node{
				Val: "b",
				Left: &lowestcommonancestor.Node{
					Val: "d",
				},
				Right: &lowestcommonancestor.Node{
					Val: "e",
					Left: &lowestcommonancestor.Node{
						Val: "g",
					},
					Right: &lowestcommonancestor.Node{
						Val: "h",
					},
				},
			},
			Right: &lowestcommonancestor.Node{
				Val: "c",
				Right: &lowestcommonancestor.Node{
					Val: "f",
				},
			},
		},
		val1:     "f",
		val2:     "c",
		expected: "c",
	},
	{
		name: "test_05",
		root: &lowestcommonancestor.Node{
			Val: "l",
			Left: &lowestcommonancestor.Node{
				Val: "m",
			},
			Right: &lowestcommonancestor.Node{
				Val: "n",
				Left: &lowestcommonancestor.Node{
					Val: "o",
					Left: &lowestcommonancestor.Node{
						Val: "q",
					},
					Right: &lowestcommonancestor.Node{
						Val: "r",
					},
				},
				Right: &lowestcommonancestor.Node{
					Val: "p",
					Left: &lowestcommonancestor.Node{
						Val: "s",
					},
					Right: &lowestcommonancestor.Node{
						Val: "t",
					},
				},
			},
		},
		val1:     "r",
		val2:     "p",
		expected: "n",
	},
	{
		name: "test_06",
		root: &lowestcommonancestor.Node{
			Val: "l",
			Left: &lowestcommonancestor.Node{
				Val: "m",
			},
			Right: &lowestcommonancestor.Node{
				Val: "n",
				Left: &lowestcommonancestor.Node{
					Val: "o",
					Left: &lowestcommonancestor.Node{
						Val: "q",
					},
					Right: &lowestcommonancestor.Node{
						Val: "r",
					},
				},
				Right: &lowestcommonancestor.Node{
					Val: "p",
					Left: &lowestcommonancestor.Node{
						Val: "s",
					},
					Right: &lowestcommonancestor.Node{
						Val: "t",
					},
				},
			},
		},
		val1:     "m",
		val2:     "o",
		expected: "l",
	},
	{
		name: "test_07",
		root: &lowestcommonancestor.Node{
			Val: "l",
			Left: &lowestcommonancestor.Node{
				Val: "m",
			},
			Right: &lowestcommonancestor.Node{
				Val: "n",
				Left: &lowestcommonancestor.Node{
					Val: "o",
					Left: &lowestcommonancestor.Node{
						Val: "q",
					},
					Right: &lowestcommonancestor.Node{
						Val: "r",
					},
				},
				Right: &lowestcommonancestor.Node{
					Val: "p",
					Left: &lowestcommonancestor.Node{
						Val: "s",
					},
					Right: &lowestcommonancestor.Node{
						Val: "t",
					},
				},
			},
		},
		val1:     "m",
		val2:     "o",
		expected: "l",
	},
	{
		name: "test_08",
		root: &lowestcommonancestor.Node{
			Val: "l",
			Left: &lowestcommonancestor.Node{
				Val: "m",
			},
			Right: &lowestcommonancestor.Node{
				Val: "n",
				Left: &lowestcommonancestor.Node{
					Val: "o",
					Left: &lowestcommonancestor.Node{
						Val: "q",
					},
					Right: &lowestcommonancestor.Node{
						Val: "r",
					},
				},
				Right: &lowestcommonancestor.Node{
					Val: "p",
					Left: &lowestcommonancestor.Node{
						Val: "s",
					},
					Right: &lowestcommonancestor.Node{
						Val: "t",
					},
				},
			},
		},
		val1:     "s",
		val2:     "p",
		expected: "p",
	},
}
