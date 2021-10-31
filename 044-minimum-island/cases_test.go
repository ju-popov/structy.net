package minimumisland_test

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
			{"W", "L", "W", "W", "W"},
			{"W", "L", "W", "W", "W"},
			{"W", "W", "W", "L", "W"},
			{"W", "W", "L", "L", "W"},
			{"L", "W", "W", "L", "L"},
			{"L", "L", "W", "W", "W"},
		},
		expected: 2,
	},
	{
		name: "test_01",
		grid: [][]string{
			{"L", "W", "W", "L", "W"},
			{"L", "W", "W", "L", "L"},
			{"W", "L", "W", "L", "W"},
			{"W", "W", "W", "W", "W"},
			{"W", "W", "L", "L", "L"},
		},
		expected: 1,
	},
	{
		name: "test_02",
		grid: [][]string{
			{"L", "L", "L"},
			{"L", "L", "L"},
			{"L", "L", "L"},
		},
		expected: 9,
	},
	{
		name: "test_03",
		grid: [][]string{
			{"W", "W"},
			{"L", "L"},
			{"W", "W"},
			{"W", "L"},
		},
		expected: 1,
	},
}
