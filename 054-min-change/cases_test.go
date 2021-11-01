package minchange_test

type testCase struct {
	name     string
	amount   int
	coins    []int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		amount:   8,
		coins:    []int{1, 5, 4, 12},
		expected: 2,
	},
	{
		name:     "test_01",
		amount:   13,
		coins:    []int{1, 9, 5, 14, 30},
		expected: 5,
	},
	{
		name:     "test_02",
		amount:   23,
		coins:    []int{2, 5, 7},
		expected: 4,
	},
	{
		name:     "test_03",
		amount:   102,
		coins:    []int{1, 5, 10, 25},
		expected: 6,
	},
	{
		name:     "test_04",
		amount:   200,
		coins:    []int{1, 5, 10, 25},
		expected: 8,
	},
	{
		name:     "test_05",
		amount:   2017,
		coins:    []int{4, 2, 10},
		expected: -1,
	},
	{
		name:     "test_06",
		amount:   271,
		coins:    []int{10, 8, 265, 24},
		expected: -1,
	},
	{
		name:     "test_07",
		amount:   0,
		coins:    []int{4, 2, 10},
		expected: 0,
	},
	{
		name:     "test_08",
		amount:   0,
		coins:    []int{},
		expected: 0,
	},
}
