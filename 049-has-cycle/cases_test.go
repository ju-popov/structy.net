package hascycle_test

type testCase struct {
	name     string
	graph    map[string][]string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		graph: map[string][]string{
			"a": {"b"},
			"b": {"c"},
			"c": {"a"},
		},
		expected: true,
	},
	{
		name: "test_01",
		graph: map[string][]string{
			"a": {"b", "c"},
			"b": {"c"},
			"c": {"d"},
			"d": {},
		},
		expected: false,
	},
	{
		name: "test_02",
		graph: map[string][]string{
			"a": {"b", "c"},
			"b": {},
			"c": {},
			"e": {"f"},
			"f": {"e"},
		},
		expected: true,
	},
	{
		name: "test_03",
		graph: map[string][]string{
			"q": {"r", "s"},
			"r": {"t", "u"},
			"s": {},
			"t": {},
			"u": {},
			"v": {"w"},
			"w": {},
			"x": {"w"},
		},
		expected: false,
	},
	{
		name: "test_04",
		graph: map[string][]string{
			"a": {"b"},
			"b": {"c"},
			"c": {"a"},
			"g": {},
		},
		expected: true,
	},
}
