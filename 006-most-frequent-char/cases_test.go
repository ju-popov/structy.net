package mostfrequentchar_test

type testCase struct {
	name     string
	input    string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		input:    "bookeeper",
		expected: "e",
	},
	{
		name:     "test_01",
		input:    "david",
		expected: "d",
	},
	{
		name:     "test_02",
		input:    "abby",
		expected: "b",
	},
	{
		name:     "test_03",
		input:    "mississippi",
		expected: "i",
	},
	{
		name:     "test_04",
		input:    "potato",
		expected: "o",
	},
	{
		name:     "test_05",
		input:    "eleventennine",
		expected: "e",
	},
	{
		name:     "test_06",
		input:    "riverbed",
		expected: "r",
	},
}
