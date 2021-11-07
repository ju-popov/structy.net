package isprime_test

type testCase struct {
	name     string
	n        int
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		n:        2,
		expected: true,
	},
	{
		name:     "test_01",
		n:        3,
		expected: true,
	},
	{
		name:     "test_02",
		n:        4,
		expected: false,
	},
	{
		name:     "test_03",
		n:        5,
		expected: true,
	},
	{
		name:     "test_04",
		n:        6,
		expected: false,
	},
	{
		name:     "test_05",
		n:        7,
		expected: true,
	},
	{
		name:     "test_06",
		n:        8,
		expected: false,
	},
	{
		name:     "test_07",
		n:        25,
		expected: false,
	},
	{
		name:     "test_08",
		n:        31,
		expected: true,
	},
	{
		name:     "test_09",
		n:        2017,
		expected: true,
	},
	{
		name:     "test_10",
		n:        2048,
		expected: false,
	},
	{
		name:     "test_11",
		n:        1,
		expected: false,
	},
}
