package compress_test

type testCase struct {
	name     string
	s        string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "ccaaatsss",
		expected: "2c3at3s",
	},
	{
		name:     "test_01",
		s:        "ssssbbz",
		expected: "4s2bz",
	},
	{
		name:     "test_02",
		s:        "ppoppppp",
		expected: "2po5p",
	},
	{
		name:     "test_03",
		s:        "nnneeeeeeeeeeeezz",
		expected: "3n12e2z",
	},
	{
		name:     "test_04",
		s:        "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
		expected: "127y",
	},
}
