package uncompress_test

type testCase struct {
	name     string
	input    string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		input:    "2c3a1t",
		expected: "ccaaat",
	},
	{
		name:     "test_01",
		input:    "4s2b",
		expected: "ssssbb",
	},
	{
		name:     "test_02",
		input:    "2p1o5p",
		expected: "ppoppppp",
	},
	{
		name:     "test_03",
		input:    "3n12e2z",
		expected: "nnneeeeeeeeeeeezz",
	},
	{
		name:     "test_04",
		input:    "127y",
		expected: "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
	},
}
