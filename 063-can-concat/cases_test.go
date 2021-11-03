package canconcat_test

type testCase struct {
	name     string
	s        string
	words    []string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "oneisnone",
		words:    []string{"one", "none", "is"},
		expected: true,
	},
	{
		name:     "test_01",
		s:        "oneisnone",
		words:    []string{"on", "e", "is"},
		expected: false,
	},
	{
		name:     "test_02",
		s:        "oneisnone",
		words:    []string{"on", "e", "is", "n"},
		expected: true,
	},
	{
		name:     "test_03",
		s:        "foodisgood",
		words:    []string{"is", "g", "ood", "f"},
		expected: true,
	},
	{
		name:     "test_04",
		s:        "santahat",
		words:    []string{"santah", "hat"},
		expected: false,
	},
	{
		name:     "test_05",
		s:        "santahat",
		words:    []string{"santah", "san", "hat", "tahat"},
		expected: true,
	},
	{
		name:     "test_06",
		s:        "rrrrrrrrrrrrrrrrrrrrrrrrrrx",
		words:    []string{"r", "rr", "rrr", "rrrr", "rrrrr", "rrrrrr"},
		expected: false,
	},
}
