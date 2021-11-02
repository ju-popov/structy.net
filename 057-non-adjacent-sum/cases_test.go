package nonadjacentsum_test

type testCase struct {
	name     string
	nums     []int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		nums:     []int{2, 4, 5, 12, 7},
		expected: 16,
	},
	{
		name:     "test_01",
		nums:     []int{7, 5, 5, 12},
		expected: 19,
	},
	{
		name:     "test_02",
		nums:     []int{7, 5, 5, 12, 17, 29},
		expected: 48,
	},
	{
		name: "test_03",
		nums: []int{
			72, 62, 10, 6, 20, 19, 42,
			46, 24, 78, 30, 41, 75, 38,
			23, 28, 66, 55, 12, 17, 9,
			12, 3, 1, 19, 30, 50, 20,
		},
		expected: 488,
	},
	{
		name: "test_04",
		nums: []int{
			72, 62, 10, 6, 20, 19, 42, 46, 24, 78,
			30, 41, 75, 38, 23, 28, 66, 55, 12, 17,
			83, 80, 56, 68, 6, 22, 56, 96, 77, 98,
			61, 20, 0, 76, 53, 74, 8, 22, 92, 37,
			30, 41, 75, 38, 23, 28, 66, 55, 12, 17,
			72, 62, 10, 6, 20, 19, 42, 46, 24, 78,
			42,
		},
		expected: 1465,
	},
}
