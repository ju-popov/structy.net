package fivesort_test

type testCase struct {
	name     string
	nums     []int64
	expected []int64
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		nums:     []int64{12, 5, 1, 5, 12, 7},
		expected: []int64{12, 7, 1, 12, 5, 5},
	},
	{
		name:     "test_01",
		nums:     []int64{12, 5, 1, 5, 12, 7},
		expected: []int64{12, 7, 1, 12, 5, 5},
	},
	{
		name:     "test_02",
		nums:     []int64{5, 5, 5, 1, 1, 1, 4},
		expected: []int64{4, 1, 1, 1, 5, 5, 5},
	},
	{
		name:     "test_03",
		nums:     []int64{5, 5, 6, 5, 5, 5, 5},
		expected: []int64{6, 5, 5, 5, 5, 5, 5},
	},
	{
		name:     "test_04",
		nums:     []int64{5, 1, 2, 5, 5, 3, 2, 5, 1, 5, 5, 5, 4, 5},
		expected: []int64{4, 1, 2, 1, 2, 3, 5, 5, 5, 5, 5, 5, 5, 5},
	},
}
