package closestcarrot_test

type testCase struct {
	name     string
	grid     [][]string
	startRow int
	startCol int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		grid: [][]string{
			{"O", "O", "O", "O", "O"},
			{"O", "X", "O", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"O", "X", "C", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"C", "O", "O", "O", "O"},
		},
		startRow: 1,
		startCol: 2,
		expected: 4,
	},
	{
		name: "test_01",
		grid: [][]string{
			{"O", "O", "O", "O", "O"},
			{"O", "X", "O", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"O", "X", "C", "O", "O"},
			{"O", "X", "X", "O", "O"},
			{"C", "O", "O", "O", "O"},
		},
		startRow: 0,
		startCol: 0,
		expected: 5,
	},
	{
		name: "test_02",
		grid: [][]string{
			{"O", "O", "X", "X", "X"},
			{"O", "X", "X", "X", "C"},
			{"O", "X", "O", "X", "X"},
			{"O", "O", "O", "O", "O"},
			{"O", "X", "X", "X", "X"},
			{"O", "O", "O", "O", "O"},
			{"O", "O", "C", "O", "O"},
			{"O", "O", "O", "O", "O"},
		},
		startRow: 3,
		startCol: 4,
		expected: 9,
	},
	{
		name: "test_03",
		grid: [][]string{
			{"O", "O", "X", "O", "O"},
			{"O", "X", "X", "X", "O"},
			{"O", "X", "C", "C", "O"},
		},
		startRow: 1,
		startCol: 4,
		expected: 2,
	},
	{
		name: "test_04",
		grid: [][]string{
			{"O", "O", "X", "O", "O"},
			{"O", "X", "X", "X", "O"},
			{"O", "X", "C", "C", "O"},
		},
		startRow: 2,
		startCol: 0,
		expected: -1,
	},
}
