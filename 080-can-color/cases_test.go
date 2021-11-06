package cancolor_test

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
			"x": {"y"},
			"y": {"x", "z"},
			"z": {"y"},
		},
		expected: true,
	},
	{
		name: "test_01",
		graph: map[string][]string{
			"q": {"r", "s"},
			"r": {"q", "s"},
			"s": {"r", "q"},
		},
		expected: false,
	},
	{
		name: "test_02",
		graph: map[string][]string{
			"a": {"b", "c", "d"},
			"b": {"a"},
			"c": {"a"},
			"d": {"a"},
		},
		expected: true,
	},
	{
		name: "test_03",
		graph: map[string][]string{
			"a": {"b", "c", "d"},
			"b": {"a"},
			"c": {"a", "d"},
			"d": {"a", "c"},
		},
		expected: false,
	},
	{
		name: "test_04",
		graph: map[string][]string{
			"h": {"i", "k"},
			"i": {"h", "j"},
			"j": {"i", "k"},
			"k": {"h", "j"},
		},
		expected: true,
	},
	{
		name:     "test_05",
		graph:    map[string][]string{},
		expected: true,
	},
	{
		name: "test_06",
		graph: map[string][]string{
			"h": {"i", "k"},
			"i": {"h", "j"},
			"j": {"i", "k"},
			"k": {"h", "j"},
			"q": {"r", "s"},
			"r": {"q", "s"},
			"s": {"r", "q"},
		},
		expected: false,
	},
}
