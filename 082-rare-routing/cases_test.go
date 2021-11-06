package rarerouting_test

type testCase struct {
	name     string
	n        int
	roads    [][2]int
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		n:    4,
		roads: [][2]int{
			{0, 1},
			{0, 2},
			{0, 3},
		},
		expected: true,
	},
	{
		name: "test_01",
		n:    4,
		roads: [][2]int{
			{0, 1},
			{0, 2},
			{0, 3},
			{3, 2},
		},
		expected: false,
	},
	{
		name: "test_02",
		n:    6,
		roads: [][2]int{
			{1, 2},
			{5, 4},
			{3, 0},
			{0, 1},
			{0, 4},
		},
		expected: true,
	},
	{
		name: "test_03",
		n:    6,
		roads: [][2]int{
			{1, 2},
			{4, 1},
			{5, 4},
			{3, 0},
			{0, 1},
			{0, 4},
		},
		expected: false,
	},
	{
		name: "test_04",
		n:    4,
		roads: [][2]int{
			{0, 1},
			{3, 2},
		},
		expected: false,
	},
}
