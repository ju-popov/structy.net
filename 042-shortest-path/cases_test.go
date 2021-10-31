package shortestpath_test

type testCase struct {
	name     string
	edges    [][2]string
	nodeA    string
	nodeB    string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		edges: [][2]string{
			{"w", "x"},
			{"x", "y"},
			{"z", "y"},
			{"z", "v"},
			{"w", "v"},
		},
		nodeA:    "w",
		nodeB:    "z",
		expected: 2,
	},
	{
		name: "test_01",
		edges: [][2]string{
			{"w", "x"},
			{"x", "y"},
			{"z", "y"},
			{"z", "v"},
			{"w", "v"},
		},
		nodeA:    "y",
		nodeB:    "x",
		expected: 1,
	},
	{
		name: "test_02",
		edges: [][2]string{
			{"a", "c"},
			{"a", "b"},
			{"c", "b"},
			{"c", "d"},
			{"b", "d"},
			{"e", "d"},
			{"g", "f"},
		},
		nodeA:    "a",
		nodeB:    "e",
		expected: 3,
	},
	{
		name: "test_03",
		edges: [][2]string{
			{"a", "c"},
			{"a", "b"},
			{"c", "b"},
			{"c", "d"},
			{"b", "d"},
			{"e", "d"},
			{"g", "f"},
		},
		nodeA:    "e",
		nodeB:    "c",
		expected: 2,
	},
	{
		name: "test_04",
		edges: [][2]string{
			{"a", "c"},
			{"a", "b"},
			{"c", "b"},
			{"c", "d"},
			{"b", "d"},
			{"e", "d"},
			{"g", "f"},
		},
		nodeA:    "b",
		nodeB:    "g",
		expected: -1,
	},
	{
		name: "test_05",
		edges: [][2]string{
			{"c", "n"},
			{"c", "e"},
			{"c", "s"},
			{"c", "w"},
			{"w", "e"},
		},
		nodeA:    "w",
		nodeB:    "e",
		expected: 1,
	},
	{
		name: "test_06",
		edges: [][2]string{
			{"c", "n"},
			{"c", "e"},
			{"c", "s"},
			{"c", "w"},
			{"w", "e"},
		},
		nodeA:    "n",
		nodeB:    "e",
		expected: 2,
	},
	{
		name: "test_07",
		edges: [][2]string{
			{"m", "n"},
			{"n", "o"},
			{"o", "p"},
			{"p", "q"},
			{"t", "o"},
			{"r", "q"},
			{"r", "s"},
		},
		nodeA:    "m",
		nodeB:    "s",
		expected: 6,
	},
}
