package mostfrequentchar_test

type testCase struct {
	name     string
	s        string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "bookeeper",
		expected: "e",
	},
	{
		name:     "test_01",
		s:        "david",
		expected: "d",
	},
	{
		name:     "test_02",
		s:        "abby",
		expected: "b",
	},
	{
		name:     "test_03",
		s:        "mississippi",
		expected: "i",
	},
	{
		name:     "test_04",
		s:        "potato",
		expected: "o",
	},
	{
		name:     "test_05",
		s:        "eleventennine",
		expected: "e",
	},
	{
		name:     "test_06",
		s:        "riverbed",
		expected: "r",
	},
}
