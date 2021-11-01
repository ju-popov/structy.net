package countpaths_test

type testCase struct {
	name     string
	grid     [][]string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		grid: [][]string{
			{"O", "O"},
			{"O", "O"},
		},
		expected: 2,
	},
	{
		name: "test_01",
		grid: [][]string{
			{"O", "O", "X"},
			{"O", "O", "O"},
			{"O", "O", "O"},
		},
		expected: 5,
	},
	{
		name: "test_02",
		grid: [][]string{
			{"O", "O", "O"},
			{"O", "O", "X"},
			{"O", "O", "O"},
		},
		expected: 3,
	},
	{
		name: "test_03",
		grid: [][]string{
			{"O", "O", "O"},
			{"O", "X", "X"},
			{"O", "O", "O"},
		},
		expected: 1,
	},
	{
		name: "test_04",
		grid: [][]string{
			{"O", "O", "X", "O", "O", "O"},
			{"O", "O", "X", "O", "O", "O"},
			{"X", "O", "X", "O", "O", "O"},
			{"X", "X", "X", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O"},
		},
		expected: 0,
	},
	{
		name: "test_05",
		grid: [][]string{
			{"O", "O", "X", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "X"},
			{"X", "O", "O", "O", "O", "O"},
			{"X", "X", "X", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O"},
		},
		expected: 42,
	},
	{
		name: "test_06",
		grid: [][]string{
			{"O", "O", "X", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "X"},
			{"X", "O", "O", "O", "O", "O"},
			{"X", "X", "X", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "X"},
		},
		expected: 0,
	},
	{
		name: "test_07",
		grid: [][]string{
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
		},
		expected: 40116600,
	},
	{
		name: "test_08",
		grid: [][]string{
			{"O", "O", "X", "X", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "X", "X", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "X", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
			{"X", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O", "O"},
			{"X", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "X", "X", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "X", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O"},
			{"X", "X", "X", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "X", "X", "O", "O", "O", "O", "X", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "X", "X", "O", "O", "O", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O"},
			{"O", "O", "O", "O", "O", "O", "O", "O", "X", "O", "O", "O", "O", "O", "O"},
		},
		expected: 3190434,
	},
}
