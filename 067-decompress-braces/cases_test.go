package decompressbraces_test

type testCase struct {
	name     string
	s        string
	expected string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "2{q}3{tu}v",
		expected: "qqtututuv",
	},
	{
		name:     "test_01",
		s:        "ch3{ao}",
		expected: "chaoaoao",
	},
	{
		name:     "test_02",
		s:        "2{y3{o}}s",
		expected: "yoooyooos",
	},
	{
		name:     "test_03",
		s:        "z3{a2{xy}b}",
		expected: "zaxyxybaxyxybaxyxyb",
	},
	{
		name:     "test_04",
		s:        "2{3{r4{e}r}io}",
		expected: "reeeerreeeerreeeerioreeeerreeeerreeeerio",
	},
	{
		name:     "test_05",
		s:        "go3{spinn2{ing}s}",
		expected: "gospinningingsspinningingsspinningings",
	},
	{
		name:     "test_06",
		s:        "2{l2{if}azu}l",
		expected: "lififazulififazul",
	},
	{
		name:     "test_07",
		s:        "3{al4{ec}2{icia}}",
		expected: "alececececiciaiciaalececececiciaiciaalececececiciaicia",
	},
}
