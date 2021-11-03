package canconcat_test

type testCase struct {
	name     string
	s        string
	words    []string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "caution",
		words:    []string{"ca", "ion", "caut", "ut"},
		expected: 2,
	},
	{
		name:     "test_01",
		s:        "caution",
		words:    []string{"ion", "caut", "caution"},
		expected: 1,
	},
	{
		name:     "test_02",
		s:        "respondorreact",
		words:    []string{"re", "or", "spond", "act", "respond"},
		expected: 4,
	},
	{
		name:     "test_03",
		s:        "simchacindy",
		words:    []string{"sim", "simcha", "acindy", "ch"},
		expected: 3,
	},
	{
		name:     "test_04",
		s:        "simchacindy",
		words:    []string{"sim", "simcha", "acindy"},
		expected: -1,
	},
	{
		name:     "test_05",
		s:        "uuuuuu",
		words:    []string{"u", "uu", "uuu", "uuuu"},
		expected: 2,
	},
	{
		name:     "test_06",
		s:        "rongbetty",
		words:    []string{"wrong", "bet"},
		expected: -1,
	},
}
