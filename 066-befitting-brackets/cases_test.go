package befittingbrackets_test

type testCase struct {
	name     string
	str      string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		str:      "(){}[](())",
		expected: true,
	},
	{
		name:     "test_01",
		str:      "({[]})",
		expected: true,
	},
	{
		name:     "test_02",
		str:      "[][}",
		expected: false,
	},
	{
		name:     "test_03",
		str:      "{[]}({}",
		expected: false,
	},
	{
		name:     "test_04",
		str:      "[]{}(}[]",
		expected: false,
	},
	{
		name:     "test_05",
		str:      "[]{}()[]",
		expected: true,
	},
	{
		name:     "test_06",
		str:      "]{}",
		expected: false,
	},
	{
		name:     "test_07",
		str:      "",
		expected: true,
	},
}
