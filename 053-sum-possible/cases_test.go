package sumpossible_test

type testCase struct {
	name     string
	amount   int
	numbers  []int
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		amount:   8,
		numbers:  []int{5, 12, 4},
		expected: true,
	},
	{
		name:     "test_01",
		amount:   15,
		numbers:  []int{6, 2, 10, 19},
		expected: false,
	},
	{
		name:     "test_02",
		amount:   13,
		numbers:  []int{6, 2, 1},
		expected: true,
	},
	{
		name:     "test_03",
		amount:   103,
		numbers:  []int{6, 20, 1},
		expected: true,
	},
	{
		name:     "test_04",
		amount:   12,
		numbers:  []int{},
		expected: false,
	},
	{
		name:     "test_05",
		amount:   12,
		numbers:  []int{12},
		expected: true,
	},
	{
		name:     "test_06",
		amount:   0,
		numbers:  []int{},
		expected: true,
	},
	{
		name:     "test_07",
		amount:   271,
		numbers:  []int{10, 8, 265, 24},
		expected: false,
	},
	{
		name:     "test_08",
		amount:   2017,
		numbers:  []int{4, 2, 10},
		expected: false,
	},
}
