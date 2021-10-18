package maxroottoleafpathsum_test

import (
	maxroottoleafpathsum "github.com/ju-popov/structy.net/029-max-root-to-leaf-path-sum"
)

type testCase struct {
	name     string
	root     *maxroottoleafpathsum.Node
	expected int64
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &maxroottoleafpathsum.Node{
			Value: 3,
			Left: &maxroottoleafpathsum.Node{
				Value: 11,
				Left: &maxroottoleafpathsum.Node{
					Value: 4,
				},
				Right: &maxroottoleafpathsum.Node{
					Value: -2,
				},
			},
			Right: &maxroottoleafpathsum.Node{
				Value: 4,
				Right: &maxroottoleafpathsum.Node{
					Value: 1,
				},
			},
		},
		expected: 18,
	},
	{
		name: "test_01",
		root: &maxroottoleafpathsum.Node{
			Value: 5,
			Left: &maxroottoleafpathsum.Node{
				Value: 11,
				Left: &maxroottoleafpathsum.Node{
					Value: 20,
				},
				Right: &maxroottoleafpathsum.Node{
					Value: 15,
					Left: &maxroottoleafpathsum.Node{
						Value: 1,
					},
					Right: &maxroottoleafpathsum.Node{
						Value: 3,
					},
				},
			},
			Right: &maxroottoleafpathsum.Node{
				Value: 54,
			},
		},
		expected: 59,
	},
	{
		name: "test_02",
		root: &maxroottoleafpathsum.Node{
			Value: -1,
			Left: &maxroottoleafpathsum.Node{
				Value: -6,
				Left: &maxroottoleafpathsum.Node{
					Value: -3,
				},
				Right: &maxroottoleafpathsum.Node{
					Value: 0,
					Left: &maxroottoleafpathsum.Node{
						Value: -1,
					},
				},
			},
			Right: &maxroottoleafpathsum.Node{
				Value: -5,
				Right: &maxroottoleafpathsum.Node{
					Value: -13,
					Right: &maxroottoleafpathsum.Node{
						Value: -2,
					},
				},
			},
		},
		expected: -8,
	},
	{
		name: "test_03",
		root: &maxroottoleafpathsum.Node{
			Value: 42,
		},
		expected: 42,
	},
}
