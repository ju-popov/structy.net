package summingsquares_test

type testCase struct {
	name     string
	n        int
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		n:        8,
		expected: 2,
	},
	{
		name:     "test_01",
		n:        9,
		expected: 1,
	},
	{
		name:     "test_02",
		n:        12,
		expected: 3,
	},
	{
		name:     "test_03",
		n:        1,
		expected: 1,
	},
	{
		name:     "test_04",
		n:        31,
		expected: 4,
	},
	{
		name:     "test_05",
		n:        50,
		expected: 2,
	},
	{
		name:     "test_06",
		n:        68,
		expected: 2,
	},
	{
		name:     "test_07",
		n:        87,
		expected: 4,
	},
}
