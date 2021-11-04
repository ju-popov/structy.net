package levelaverages_test

import (
	levelaverages "github.com/ju-popov/structy.net/036-level-averages"
)

type testCase struct {
	name     string
	root     *levelaverages.Node
	expected []float64
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		root: &levelaverages.Node{
			Val: 3,
			Left: &levelaverages.Node{
				Val: 11,
				Left: &levelaverages.Node{
					Val: 4,
				},
				Right: &levelaverages.Node{
					Val: -2,
				},
			},
			Right: &levelaverages.Node{
				Val: 4,
				Right: &levelaverages.Node{
					Val: 1,
				},
			},
		},
		expected: []float64{3.0, (11.0 + 4.0) / 2.0, (4.0 - 2.0 + 1.0) / 3.0},
	},
	{
		name: "test_01",
		root: &levelaverages.Node{
			Val: 5,
			Left: &levelaverages.Node{
				Val: 11,
				Left: &levelaverages.Node{
					Val: 20,
				},
				Right: &levelaverages.Node{
					Val: 15,
					Left: &levelaverages.Node{
						Val: 1,
					},
					Right: &levelaverages.Node{
						Val: 3,
					},
				},
			},
			Right: &levelaverages.Node{
				Val: 54,
			},
		},
		expected: []float64{5.0, (11.0 + 54.0) / 2.0, (20.0 + 15.0) / 2.0, (1.0 + 3.0) / 2},
	},
	{
		name: "test_02",
		root: &levelaverages.Node{
			Val: -1,
			Left: &levelaverages.Node{
				Val: -6,
				Left: &levelaverages.Node{
					Val: -3,
				},
				Right: &levelaverages.Node{
					Val: 0,
					Left: &levelaverages.Node{
						Val: -1,
					},
				},
			},
			Right: &levelaverages.Node{
				Val: -5,
				Right: &levelaverages.Node{
					Val: 45,
					Right: &levelaverages.Node{
						Val: -2,
					},
				},
			},
		},
		expected: []float64{-1.0, (-6.0 - 5.0) / 2.0, (-3.0 + 0.0 + 45.0) / 3.0, (-1.0 - 2.0) / 2},
	},
	{
		name: "test_03",
		root: &levelaverages.Node{
			Val: 13,
			Left: &levelaverages.Node{
				Val: 4,
				Right: &levelaverages.Node{
					Val: 9,
					Left: &levelaverages.Node{
						Val: 2,
						Left: &levelaverages.Node{
							Val: 42,
						},
					},
				},
			},
			Right: &levelaverages.Node{
				Val: 2,
			},
		},
		expected: []float64{13.0, (4.0 + 2.0) / 2.0, 9.0, 2.0, 42.0},
	},
	{
		name:     "test_04",
		root:     nil,
		expected: []float64{},
	},
}
