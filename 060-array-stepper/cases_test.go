package countingchange_test

type testCase struct {
	name     string
	nums     []int
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		nums:     []int{2, 4, 2, 0, 0, 1},
		expected: true,
	},
	{
		name:     "test_01",
		nums:     []int{2, 3, 2, 0, 0, 1},
		expected: false,
	},
	{
		name:     "test_02",
		nums:     []int{3, 1, 3, 1, 0, 1},
		expected: true,
	},
	{
		name:     "test_03",
		nums:     []int{4, 1, 5, 1, 1, 1, 0, 4},
		expected: true,
	},
	{
		name:     "test_04",
		nums:     []int{4, 1, 2, 1, 1, 1, 0, 4},
		expected: false,
	},
	{
		name:     "test_05",
		nums:     []int{1, 1, 1, 1, 1, 0},
		expected: true,
	},
	{
		name:     "test_06",
		nums:     []int{1, 1, 1, 1, 0, 0},
		expected: false,
	},
	{
		name: "test_07",
		nums: []int{
			31, 30, 29, 28, 27,
			26, 25, 24, 23, 22,
			21, 20, 19, 18, 17,
			16, 15, 14, 13, 12,
			11, 10, 9, 8, 7, 6,
			5, 3, 2, 1, 0, 0, 0,
		},
		expected: false,
	},
}
