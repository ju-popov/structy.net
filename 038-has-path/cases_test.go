package haspath_test

type testCase struct {
	name     string
	graph    map[string][]string
	src      string
	dst      string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		graph: map[string][]string{
			"f": {"g", "i"},
			"g": {"h"},
			"h": {},
			"i": {"g", "k"},
			"j": {"i"},
			"k": {},
		},
		src:      "f",
		dst:      "k",
		expected: true,
	},
	{
		name: "test_01",
		graph: map[string][]string{
			"f": {"g", "i"},
			"g": {"h"},
			"h": {},
			"i": {"g", "k"},
			"j": {"i"},
			"k": {},
		},
		src:      "f",
		dst:      "j",
		expected: false,
	},
	{
		name: "test_02",
		graph: map[string][]string{
			"f": {"g", "i"},
			"g": {"h"},
			"h": {},
			"i": {"g", "k"},
			"j": {"i"},
			"k": {},
		},
		src:      "i",
		dst:      "h",
		expected: true,
	},
	{
		name: "test_03",
		graph: map[string][]string{
			"v": {"x", "w"},
			"w": {},
			"x": {},
			"y": {"z"},
			"z": {},
		},
		src:      "v",
		dst:      "w",
		expected: true,
	},
	{
		name: "test_04",
		graph: map[string][]string{
			"v": {"x", "w"},
			"w": {},
			"x": {},
			"y": {"z"},
			"z": {},
		},
		src:      "v",
		dst:      "z",
		expected: false,
	},
}
