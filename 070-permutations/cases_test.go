package permutations_test

type testCase struct {
	name     string
	items    []interface{}
	expected [][]interface{}
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:  "test_00",
		items: []interface{}{"a", "b", "c"},
		expected: [][]interface{}{
			{"a", "b", "c"},
			{"b", "a", "c"},
			{"b", "c", "a"},
			{"a", "c", "b"},
			{"c", "a", "b"},
			{"c", "b", "a"},
		},
	},
	{
		name:  "test_01",
		items: []interface{}{"red", "blue"},
		expected: [][]interface{}{
			{"red", "blue"},
			{"blue", "red"},
		},
	},
	{
		name:  "test_02",
		items: []interface{}{8, 2, 1, 4},
		expected: [][]interface{}{
			{8, 2, 1, 4},
			{2, 8, 1, 4},
			{2, 1, 8, 4},
			{2, 1, 4, 8},
			{8, 1, 2, 4},
			{1, 8, 2, 4},
			{1, 2, 8, 4},
			{1, 2, 4, 8},
			{8, 1, 4, 2},
			{1, 8, 4, 2},
			{1, 4, 8, 2},
			{1, 4, 2, 8},
			{8, 2, 4, 1},
			{2, 8, 4, 1},
			{2, 4, 8, 1},
			{2, 4, 1, 8},
			{8, 4, 2, 1},
			{4, 8, 2, 1},
			{4, 2, 8, 1},
			{4, 2, 1, 8},
			{8, 4, 1, 2},
			{4, 8, 1, 2},
			{4, 1, 8, 2},
			{4, 1, 2, 8},
		},
	},
	{
		name:  "test_03",
		items: []interface{}{},
		expected: [][]interface{}{
			{},
		},
	},
}
