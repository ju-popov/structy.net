package uncompress_test

type testCase struct {
	name     string
	s        string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "2c3a1t",
		expected: "ccaaat",
	},
	{
		name:     "test_01",
		s:        "4s2b",
		expected: "ssssbb",
	},
	{
		name:     "test_02",
		s:        "2p1o5p",
		expected: "ppoppppp",
	},
	{
		name:     "test_03",
		s:        "3n12e2z",
		expected: "nnneeeeeeeeeeeezz",
	},
	{
		name:     "test_04",
		s:        "127y",
		expected: "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
	},
}
