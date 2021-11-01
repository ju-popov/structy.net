package tribonacci_test

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
		expected: 0,
	},
	{
		name:     "test_02",
		n:        2,
		expected: 1,
	},
	{
		name:     "test_03",
		n:        5,
		expected: 4,
	},
	{
		name:     "test_04",
		n:        7,
		expected: 13,
	},
	{
		name:     "test_05",
		n:        14,
		expected: 927,
	},
	{
		name:     "test_06",
		n:        20,
		expected: 35890,
	},
	{
		name:     "test_07",
		n:        37,
		expected: 1132436852,
	},
}
