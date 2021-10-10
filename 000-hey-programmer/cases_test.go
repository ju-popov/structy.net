package heyprogrammer_test

type testCase struct {
	name     string
	input    string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		input:    "alvin",
		expected: "hey alvin",
	},
	{
		name:     "test_01",
		input:    "jason",
		expected: "hey jason",
	},
	{
		name:     "test_02",
		input:    "how now brown cow",
		expected: "hey how now brown cow",
	},
}
