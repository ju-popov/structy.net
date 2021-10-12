package anagrams_test

type testCase struct {
	name     string
	s1       string
	s2       string
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s1:       "restful",
		s2:       "fluster",
		expected: true,
	},
	{
		name:     "test_01",
		s1:       "cats",
		s2:       "fluster",
		expected: false,
	},
	{
		name:     "test_02",
		s1:       "monkeyswrite",
		s2:       "newyorktimes",
		expected: true,
	},
	{
		name:     "test_03",
		s1:       "paper",
		s2:       "reapa",
		expected: false,
	},
	{
		name:     "test_04",
		s1:       "elbow",
		s2:       "below",
		expected: true,
	},
	{
		name:     "test_05",
		s1:       "tax",
		s2:       "taxi",
		expected: false,
	},
	{
		name:     "test_06",
		s1:       "taxi",
		s2:       "tax",
		expected: false,
	},
	{
		name:     "test_07",
		s1:       "night",
		s2:       "thing",
		expected: true,
	},
	{
		name:     "test_08",
		s1:       "abbc",
		s2:       "aabc",
		expected: false,
	},
}
