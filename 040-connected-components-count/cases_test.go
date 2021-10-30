package connectedcomponentscount_test

type testCase struct {
	name     string
	graph    map[int][]int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		graph: map[int][]int{
			0: {8, 1, 5},
			1: {0},
			5: {0, 8},
			8: {0, 5},
			2: {3, 4},
			3: {2, 4},
			4: {3, 2},
		},
		expected: 2,
	},
	{
		name: "test_01",
		graph: map[int][]int{
			1: {2},
			2: {1, 8},
			6: {7},
			9: {8},
			7: {6, 8},
			8: {9, 7, 2},
		},
		expected: 1,
	},
	{
		name: "test_02",
		graph: map[int][]int{
			3: {},
			4: {6},
			6: {4, 5, 7, 8},
			8: {6},
			7: {6},
			5: {6},
			1: {2},
			2: {1},
		},
		expected: 3,
	},
	{
		name:     "test_03",
		graph:    map[int][]int{},
		expected: 0,
	},
	{
		name: "test_04",
		graph: map[int][]int{
			0: {4, 7},
			1: {},
			2: {},
			3: {6},
			4: {0},
			6: {3},
			7: {0},
			8: {},
		},
		expected: 5,
	},
}
