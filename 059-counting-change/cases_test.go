package countingchange_test

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
		amount:   4,
		coins:    []int{1, 2, 3},
		expected: 4,
	},
	{
		name:     "test_01",
		amount:   8,
		coins:    []int{1, 2, 3},
		expected: 10,
	},
	{
		name:     "test_02",
		amount:   24,
		coins:    []int{5, 7, 3},
		expected: 5,
	},
	{
		name:     "test_03",
		amount:   13,
		coins:    []int{2, 6, 12, 10},
		expected: 0,
	},
	{
		name:     "test_04",
		amount:   512,
		coins:    []int{1, 5, 10, 25},
		expected: 20119,
	},
	{
		name:     "test_05",
		amount:   1000,
		coins:    []int{1, 5, 10, 25},
		expected: 142511,
	},
	{
		name:     "test_06",
		amount:   240,
		coins:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		expected: 1525987916,
	},
}
