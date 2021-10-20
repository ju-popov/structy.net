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
			Value: 3,
			Left: &levelaverages.Node{
				Value: 11,
				Left: &levelaverages.Node{
					Value: 4,
				},
				Right: &levelaverages.Node{
					Value: -2,
				},
			},
			Right: &levelaverages.Node{
				Value: 4,
				Right: &levelaverages.Node{
					Value: 1,
				},
			},
		},
		expected: []float64{3.0, (11.0+4.0)/2.0, (4.0-2.0+1.0)/3.0},
	},
	{
		name: "test_01",
		root: &levelaverages.Node{
			Value: 5,
			Left: &levelaverages.Node{
				Value: 11,
				Left: &levelaverages.Node{
					Value: 20,
				},
				Right: &levelaverages.Node{
					Value: 15,
					Left: &levelaverages.Node{
						Value: 1,
					},
					Right: &levelaverages.Node{
						Value: 3,
					},
				},
			},
			Right: &levelaverages.Node{
				Value: 54,
			},
		},
		expected: []float64{5.0, (11.0+54.0)/2.0, (20.0+15.0)/2.0, (1.0+3.0)/2},
	},
	{
		name: "test_02",
		root: &levelaverages.Node{
			Value: -1,
			Left: &levelaverages.Node{
				Value: -6,
				Left: &levelaverages.Node{
					Value: -3,
				},
				Right: &levelaverages.Node{
					Value: 0,
					Left: &levelaverages.Node{
						Value: -1,
					},
				},
			},
			Right: &levelaverages.Node{
				Value: -5,
				Right: &levelaverages.Node{
					Value: 45,
					Right: &levelaverages.Node{
						Value: -2,
					},
				},
			},
		},
		expected: []float64{-1.0, (-6.0-5.0)/2.0, (-3.0+0.0+45.0)/3.0, (-1.0-2.0)/2},
	},
	{
		name: "test_03",
		root: &levelaverages.Node{
			Value: 13,
			Left: &levelaverages.Node{
				Value: 4,
				Right: &levelaverages.Node{
					Value: 9,
					Left: &levelaverages.Node{
						Value: 2,
						Left: &levelaverages.Node{
							Value: 42,
						},
					},
				},
			},
			Right: &levelaverages.Node{
				Value: 2,
			},
		},
		expected: []float64{13.0, (4.0+2.0)/2.0, 9.0, 2.0, 42.0},
	},
	{
		name: "test_04",
		root: nil,
		expected: []float64{},
	},
}
