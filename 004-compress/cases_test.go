package compress_test

type testCase struct {
	name     string
	input    string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		input:    "ccaaatsss",
		expected: "2c3at3s",
	},
	{
		name:     "test_01",
		input:    "ssssbbz",
		expected: "4s2bz",
	},
	{
		name:     "test_02",
		input:    "ppoppppp",
		expected: "2po5p",
	},
	{
		name:     "test_03",
		input:    "nnneeeeeeeeeeeezz",
		expected: "3n12e2z",
	},
	{
		name:     "test_04",
		input:    "yyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyyy",
		expected: "127y",
	},
}
