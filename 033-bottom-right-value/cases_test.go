package bottomrightvalue_test

import (
	bottomrightvalue "github.com/ju-popov/structy.net/033-bottom-right-value"
)

type testCase struct {
	name     string
	root     *bottomrightvalue.Node
	expected interface{}
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &bottomrightvalue.Node{
			Val: 3,
			Left: &bottomrightvalue.Node{
				Val: 11,
				Left: &bottomrightvalue.Node{
					Val: 4,
				},
				Right: &bottomrightvalue.Node{
					Val: -2,
				},
			},
			Right: &bottomrightvalue.Node{
				Val: 10,
				Right: &bottomrightvalue.Node{
					Val: 1,
				},
			},
		},
		expected: 1,
	},
	{
		name: "test_01",
		root: &bottomrightvalue.Node{
			Val: -1,
			Left: &bottomrightvalue.Node{
				Val: -6,
				Left: &bottomrightvalue.Node{
					Val: -3,
				},
				Right: &bottomrightvalue.Node{
					Val: -4,
					Left: &bottomrightvalue.Node{
						Val: -2,
					},
					Right: &bottomrightvalue.Node{
						Val: 6,
					},
				},
			},
			Right: &bottomrightvalue.Node{
				Val: -5,
				Right: &bottomrightvalue.Node{
					Val: -13,
				},
			},
		},
		expected: 6,
	},
	{
		name: "test_02",
		root: &bottomrightvalue.Node{
			Val: -1,
			Left: &bottomrightvalue.Node{
				Val: -6,
				Left: &bottomrightvalue.Node{
					Val: -3,
				},
				Right: &bottomrightvalue.Node{
					Val: -4,
					Left: &bottomrightvalue.Node{
						Val: -2,
					},
					Right: &bottomrightvalue.Node{
						Val: 6,
					},
				},
			},
			Right: &bottomrightvalue.Node{
				Val: -5,
				Right: &bottomrightvalue.Node{
					Val: -13,
					Left: &bottomrightvalue.Node{
						Val: 7,
					},
				},
			},
		},
		expected: 7,
	},
	{
		name: "test_03",
		root: &bottomrightvalue.Node{
			Val: "a",
			Left: &bottomrightvalue.Node{
				Val: "b",
				Right: &bottomrightvalue.Node{
					Val: "d",
					Left: &bottomrightvalue.Node{
						Val: "e",
						Left: &bottomrightvalue.Node{
							Val: "f",
						},
					},
				},
			},
			Right: &bottomrightvalue.Node{
				Val: "c",
			},
		},
		expected: "f",
	},
	{
		name: "test_04",
		root: &bottomrightvalue.Node{
			Val: 42,
		},
		expected: 42,
	},
}
