package maxpalinsubsequence_test

type testCase struct {
	name     string
	str      string
	expected int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		str:      "luwxult",
		expected: 5,
	},
	{
		name:     "test_01",
		str:      "xyzaxxzy",
		expected: 6,
	},
	{
		name:     "test_02",
		str:      "lol",
		expected: 3,
	},
	{
		name:     "test_03",
		str:      "boabcdefop",
		expected: 3,
	},
	{
		name:     "test_04",
		str:      "z",
		expected: 1,
	},
	{
		name:     "test_05",
		str:      "chartreusepugvicefree",
		expected: 7,
	},
	{
		name:     "test_06",
		str:      "qwueoiuahsdjnweuueueunasdnmnqweuzqwerty",
		expected: 15,
	},
	{
		name:     "test_07",
		str:      "enamelpinportlandtildecoldpressedironyflannelsemioticsedisonbulbfashionaxe",
		expected: 31,
	},
}
