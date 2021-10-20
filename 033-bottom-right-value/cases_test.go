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
			Value: 3,
			Left: &bottomrightvalue.Node{
				Value: 11,
				Left: &bottomrightvalue.Node{
					Value: 4,
				},
				Right: &bottomrightvalue.Node{
					Value: -2,
				},
			},
			Right: &bottomrightvalue.Node{
				Value: 10,
				Right: &bottomrightvalue.Node{
					Value: 1,
				},
			},
		},
		expected: 1,
	},
	{
		name: "test_01",
		root: &bottomrightvalue.Node{
			Value: -1,
			Left: &bottomrightvalue.Node{
				Value: -6,
				Left: &bottomrightvalue.Node{
					Value: -3,
				},
				Right: &bottomrightvalue.Node{
					Value: -4,
					Left: &bottomrightvalue.Node{
						Value: -2,
					},
					Right: &bottomrightvalue.Node{
						Value: 6,
					},
				},
			},
			Right: &bottomrightvalue.Node{
				Value: -5,
				Right: &bottomrightvalue.Node{
					Value: -13,
				},
			},
		},
		expected: 6,
	},
	{
		name: "test_02",
		root: &bottomrightvalue.Node{
			Value: -1,
			Left: &bottomrightvalue.Node{
				Value: -6,
				Left: &bottomrightvalue.Node{
					Value: -3,
				},
				Right: &bottomrightvalue.Node{
					Value: -4,
					Left: &bottomrightvalue.Node{
						Value: -2,
					},
					Right: &bottomrightvalue.Node{
						Value: 6,
					},
				},
			},
			Right: &bottomrightvalue.Node{
				Value: -5,
				Right: &bottomrightvalue.Node{
					Value: -13,
					Left: &bottomrightvalue.Node{
						Value: 7,
					},
				},
			},
		},
		expected: 7,
	},
	{
		name: "test_03",
		root: &bottomrightvalue.Node{
			Value: "a",
			Left: &bottomrightvalue.Node{
				Value: "b",
				Right: &bottomrightvalue.Node{
					Value: "d",
					Left: &bottomrightvalue.Node{
						Value: "e",
						Left: &bottomrightvalue.Node{
							Value: "f",
						},
					},
				},
			},
			Right: &bottomrightvalue.Node{
				Value: "c",
			},
		},
		expected: "f",
	},
	{
		name: "test_04",
		root: &bottomrightvalue.Node{
			Value: 42,
		},
		expected: 42,
	},
}
