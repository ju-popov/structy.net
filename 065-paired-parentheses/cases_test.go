package pairedparentheses_test

type testCase struct {
	name     string
	str      string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		str:      "(david)((abby))",
		expected: true,
	},
	{
		name:     "test_01",
		str:      "()rose(jeff",
		expected: false,
	},
	{
		name:     "test_02",
		str:      ")(",
		expected: false,
	},
	{
		name:     "test_03",
		str:      "()",
		expected: true,
	},
	{
		name:     "test_04",
		str:      "(((potato())))",
		expected: true,
	},
	{
		name:     "test_05",
		str:      "(())(water)()",
		expected: true,
	},
	{
		name:     "test_06",
		str:      "(())(water()()",
		expected: false,
	},
	{
		name:     "test_07",
		str:      "",
		expected: true,
	},
}
