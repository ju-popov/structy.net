package maxroottoleafpathsum_test

import (
	maxroottoleafpathsum "github.com/ju-popov/structy.net/029-max-root-to-leaf-path-sum"
)

type testCase struct {
	name     string
	root     *maxroottoleafpathsum.Node
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &maxroottoleafpathsum.Node{
			Val: 3,
			Left: &maxroottoleafpathsum.Node{
				Val: 11,
				Left: &maxroottoleafpathsum.Node{
					Val: 4,
				},
				Right: &maxroottoleafpathsum.Node{
					Val: -2,
				},
			},
			Right: &maxroottoleafpathsum.Node{
				Val: 4,
				Right: &maxroottoleafpathsum.Node{
					Val: 1,
				},
			},
		},
		expected: 18,
	},
	{
		name: "test_01",
		root: &maxroottoleafpathsum.Node{
			Val: 5,
			Left: &maxroottoleafpathsum.Node{
				Val: 11,
				Left: &maxroottoleafpathsum.Node{
					Val: 20,
				},
				Right: &maxroottoleafpathsum.Node{
					Val: 15,
					Left: &maxroottoleafpathsum.Node{
						Val: 1,
					},
					Right: &maxroottoleafpathsum.Node{
						Val: 3,
					},
				},
			},
			Right: &maxroottoleafpathsum.Node{
				Val: 54,
			},
		},
		expected: 59,
	},
	{
		name: "test_02",
		root: &maxroottoleafpathsum.Node{
			Val: -1,
			Left: &maxroottoleafpathsum.Node{
				Val: -6,
				Left: &maxroottoleafpathsum.Node{
					Val: -3,
				},
				Right: &maxroottoleafpathsum.Node{
					Val: 0,
					Left: &maxroottoleafpathsum.Node{
						Val: -1,
					},
				},
			},
			Right: &maxroottoleafpathsum.Node{
				Val: -5,
				Right: &maxroottoleafpathsum.Node{
					Val: -13,
					Right: &maxroottoleafpathsum.Node{
						Val: -2,
					},
				},
			},
		},
		expected: -8,
	},
	{
		name: "test_03",
		root: &maxroottoleafpathsum.Node{
			Val: 42,
		},
		expected: 42,
	},
}
