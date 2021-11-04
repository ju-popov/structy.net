package pairsum_test

type testCase struct {
	name      string
	numbers   []int
	targetSum int
	expected  [2]int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:      "test_00",
		numbers:   []int{3, 2, 5, 4, 1},
		targetSum: 8,
		expected:  [2]int{0, 2},
	},
	{
		name:      "test_01",
		numbers:   []int{4, 7, 9, 2, 5, 1},
		targetSum: 5,
		expected:  [2]int{0, 5},
	},
	{
		name:      "test_02",
		numbers:   []int{4, 7, 9, 2, 5, 1},
		targetSum: 3,
		expected:  [2]int{3, 5},
	},
	{
		name:      "test_03",
		numbers:   []int{1, 6, 7, 2},
		targetSum: 13,
		expected:  [2]int{1, 2},
	},
	{
		name:      "test_04",
		numbers:   []int{9, 9},
		targetSum: 18,
		expected:  [2]int{0, 1},
	},
	{
		name:      "test_05",
		numbers:   []int{6, 4, 2, 8},
		targetSum: 12,
		expected:  [2]int{1, 3},
	},
}
