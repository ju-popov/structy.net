package heyprogrammer_test

type testCase struct {
	name     string
	s        string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "alvin",
		expected: "hey alvin",
	},
	{
		name:     "test_01",
		s:        "jason",
		expected: "hey jason",
	},
	{
		name:     "test_02",
		s:        "how now brown cow",
		expected: "hey how now brown cow",
	},
}
