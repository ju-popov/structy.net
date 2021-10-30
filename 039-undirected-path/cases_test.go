package undirectedpath_test

type testCase struct {
	name     string
	edges    [][2]string
	nodeA    string
	nodeB    string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		edges: [][2]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		nodeA:    "j",
		nodeB:    "m",
		expected: true,
	},
	{
		name: "test_01",
		edges: [][2]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		nodeA:    "m",
		nodeB:    "j",
		expected: true,
	},
	{
		name: "test_02",
		edges: [][2]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		nodeA:    "l",
		nodeB:    "j",
		expected: true,
	},
	{
		name: "test_03",
		edges: [][2]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		nodeA:    "k",
		nodeB:    "o",
		expected: false,
	},
	{
		name: "test_04",
		edges: [][2]string{
			{"i", "j"},
			{"k", "i"},
			{"m", "k"},
			{"k", "l"},
			{"o", "n"},
		},
		nodeA:    "i",
		nodeB:    "o",
		expected: false,
	},
	{
		name: "test_05",
		edges: [][2]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		nodeA:    "a",
		nodeB:    "b",
		expected: true,
	},
	{
		name: "test_06",
		edges: [][2]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		nodeA:    "a",
		nodeB:    "c",
		expected: true,
	},
	{
		name: "test_07",
		edges: [][2]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		nodeA:    "r",
		nodeB:    "t",
		expected: true,
	},
	{
		name: "test_08",
		edges: [][2]string{
			{"b", "a"},
			{"c", "a"},
			{"b", "c"},
			{"q", "r"},
			{"q", "s"},
			{"q", "u"},
			{"q", "t"},
		},
		nodeA:    "r",
		nodeB:    "b",
		expected: false,
	},
}
