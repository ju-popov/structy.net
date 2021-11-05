package subsets_test

type testCase struct {
	name     string
	elements []string
	expected [][]string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		elements: []string{"a", "b"},
		expected: [][]string{
			{},
			{"b"},
			{"a"},
			{"a", "b"},
		},
	},
	{
		name:     "test_01",
		elements: []string{"a", "b", "c"},
		expected: [][]string{
			{},
			{"c"},
			{"b"},
			{"b", "c"},
			{"a"},
			{"a", "c"},
			{"a", "b"},
			{"a", "b", "c"},
		},
	},
	{
		name:     "test_02",
		elements: []string{"x"},
		expected: [][]string{
			{},
			{"x"},
		},
	},
	{
		name:     "test_03",
		elements: []string{},
		expected: [][]string{
			{},
		},
	},
	{
		name:     "test_04",
		elements: []string{"q", "r", "s", "t"},
		expected: [][]string{
			{},
			{"t"},
			{"s"},
			{"s", "t"},
			{"r"},
			{"r", "t"},
			{"r", "s"},
			{"r", "s", "t"},
			{"q"},
			{"q", "t"},
			{"q", "s"},
			{"q", "s", "t"},
			{"q", "r"},
			{"q", "r", "t"},
			{"q", "r", "s"},
			{"q", "r", "s", "t"},
		},
	},
}
