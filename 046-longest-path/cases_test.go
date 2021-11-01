package longestpath_test

type testCase struct {
	name     string
	graph    map[string][]string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		graph: map[string][]string{
			"a": {"c", "b"},
			"b": {"c"},
			"c": {},
		},
		expected: 2,
	},
	{
		name: "test_01",
		graph: map[string][]string{
			"a": {"c", "b"},
			"b": {"c"},
			"c": {},
			"q": {"r"},
			"r": {"s", "u", "t"},
			"s": {"t"},
			"t": {"u"},
			"u": {},
		},
		expected: 4,
	},
	{
		name: "test_02",
		graph: map[string][]string{
			"h": {"i", "j", "k"},
			"g": {"h"},
			"i": {},
			"j": {},
			"k": {},
			"x": {"y"},
			"y": {},
		},
		expected: 2,
	},
	{
		name: "test_03",
		graph: map[string][]string{
			"a": {"b"},
			"b": {"c"},
			"c": {},
			"e": {"f"},
			"f": {"g"},
			"g": {"h"},
			"h": {},
		},
		expected: 3,
	},
}
