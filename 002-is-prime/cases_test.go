package isprime_test

type testCase struct {
	name     string
	input    int64
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		input:    2,
		expected: true,
	},
	{
		name:     "test_01",
		input:    3,
		expected: true,
	},
	{
		name:     "test_02",
		input:    4,
		expected: false,
	},
	{
		name:     "test_03",
		input:    5,
		expected: true,
	},
	{
		name:     "test_04",
		input:    6,
		expected: false,
	},
	{
		name:     "test_05",
		input:    7,
		expected: true,
	},
	{
		name:     "test_06",
		input:    8,
		expected: false,
	},
	{
		name:     "test_07",
		input:    25,
		expected: false,
	},
	{
		name:     "test_08",
		input:    31,
		expected: true,
	},
	{
		name:     "test_09",
		input:    2017,
		expected: true,
	},
	{
		name:     "test_10",
		input:    2048,
		expected: false,
	},
	{
		name:     "test_11",
		input:    1,
		expected: false,
	},
}
