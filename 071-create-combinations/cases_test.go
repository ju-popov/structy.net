package createcombinations_test

type testCase struct {
	name     string
	items    []interface{}
	k        int
	expected [][]interface{}
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:  "test_00",
		items: []interface{}{"a", "b", "c"},
		k:     2,
		expected: [][]interface{}{
			{"a", "b"},
			{"a", "c"},
			{"b", "c"},
		},
	},
	{
		name:  "test_01",
		items: []interface{}{"q", "r", "s", "t"},
		k:     2,
		expected: [][]interface{}{
			{"q", "r"},
			{"q", "s"},
			{"q", "t"},
			{"r", "s"},
			{"r", "t"},
			{"s", "t"},
		},
	},
	{
		name:  "test_02",
		items: []interface{}{"q", "r", "s", "t"},
		k:     3,
		expected: [][]interface{}{
			{"q", "r", "s"},
			{"q", "r", "t"},
			{"q", "s", "t"},
			{"r", "s", "t"},
		},
	},
	{
		name:  "test_03",
		items: []interface{}{1, 28, 94},
		k:     3,
		expected: [][]interface{}{
			{1, 28, 94},
		},
	},
}
