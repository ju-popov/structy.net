package tolerantteams_test

type testCase struct {
	name      string
	rivalries [][2]string
	expected  bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name: "test_00",
		rivalries: [][2]string{
			{"philip", "seb"},
			{"raj", "nader"},
		},
		expected: true,
	},
	{
		name: "test_01",
		rivalries: [][2]string{
			{"philip", "seb"},
			{"raj", "nader"},
			{"raj", "philip"},
			{"seb", "raj"},
		},
		expected: false,
	},
	{
		name: "test_02",
		rivalries: [][2]string{
			{"cindy", "anj"},
			{"alex", "matt"},
			{"alex", "cindy"},
			{"anj", "matt"},
			{"brando", "matt"},
		},
		expected: true,
	},
	{
		name: "test_03",
		rivalries: [][2]string{
			{"alex", "anj"},
			{"alex", "matt"},
			{"alex", "cindy"},
			{"anj", "matt"},
			{"brando", "matt"},
		},
		expected: false,
	},
	{
		name: "test_04",
		rivalries: [][2]string{
			{"alan", "jj"},
			{"betty", "richard"},
			{"jj", "simcha"},
			{"richard", "christine"},
		},
		expected: true,
	},
	{
		name: "test_05",
		rivalries: [][2]string{
			{"alan", "jj"},
			{"betty", "richard"},
			{"jj", "simcha"},
			{"richard", "christine"},
		},
		expected: true,
	},
	{
		name: "test_06",
		rivalries: [][2]string{
			{"alan", "jj"},
			{"jj", "richard"},
			{"betty", "richard"},
			{"jj", "simcha"},
			{"richard", "christine"},
		},
		expected: true,
	},
	{
		name: "test_07",
		rivalries: [][2]string{
			{"alan", "jj"},
			{"betty", "richard"},
			{"betty", "christine"},
			{"jj", "simcha"},
			{"richard", "christine"},
		},
		expected: false,
	},
}
