package nestingscore_test

type testCase struct {
	name     string
	str      string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		str:      "[]",
		expected: 1,
	},
	{
		name:     "test_01",
		str:      "[][][]",
		expected: 3,
	},
	{
		name:     "test_02",
		str:      "[[]]",
		expected: 2,
	},
	{
		name:     "test_03",
		str:      "[[][]]",
		expected: 4,
	},
	{
		name:     "test_04",
		str:      "[[][][]]",
		expected: 6,
	},
	{
		name:     "test_05",
		str:      "[[][]][]",
		expected: 5,
	},
	{
		name:     "test_06",
		str:      "[][[][]][[]]",
		expected: 7,
	},
	{
		name:     "test_07",
		str:      "[[[[[[[][]]]]]]][]",
		expected: 129,
	},
}
