package overlapsubsequence_test

type testCase struct {
	name     string
	str1     string
	str2     string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		str1:     "dogs",
		str2:     "daogt",
		expected: 3,
	},
	{
		name:     "test_01",
		str1:     "xcyats",
		str2:     "criaotsi",
		expected: 4,
	},
	{
		name:     "test_02",
		str1:     "xfeqortsver",
		str2:     "feeeuavoeqr",
		expected: 5,
	},
	{
		name:     "test_03",
		str1:     "kinfolklivemustache",
		str2:     "bespokekinfolksnackwave",
		expected: 11,
	},
	{
		name:     "test_04",
		str1:     "mumblecorebeardleggingsauthenticunicorn",
		str2:     "succulentspughumblemeditationlocavore",
		expected: 15,
	},
}
