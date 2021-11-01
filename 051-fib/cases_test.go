package fib_test

type testCase struct {
	name     string
	n        int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		n:        0,
		expected: 0,
	},
	{
		name:     "test_01",
		n:        1,
		expected: 1,
	},
	{
		name:     "test_02",
		n:        2,
		expected: 1,
	},
	{
		name:     "test_03",
		n:        3,
		expected: 2,
	},
	{
		name:     "test_04",
		n:        4,
		expected: 3,
	},
	{
		name:     "test_05",
		n:        5,
		expected: 5,
	},
	{
		name:     "test_06",
		n:        35,
		expected: 9227465,
	},
	{
		name:     "test_07",
		n:        46,
		expected: 1836311903,
	},
}
