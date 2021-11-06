package parentheticalpossibilities_test

type testCase struct {
	name     string
	s        string
	expected []string
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		s:        "x(mn)yz",
		expected: []string{"xmyz", "xnyz"},
	},
	{
		name:     "test_01",
		s:        "(qr)ab(stu)c",
		expected: []string{"qabsc", "qabtc", "qabuc", "rabsc", "rabtc", "rabuc"},
	},
	{
		name:     "test_02",
		s:        "taco",
		expected: []string{"taco"},
	},
	{
		name:     "test_03",
		s:        "",
		expected: []string{""},
	},
	{
		name: "test_04",
		s:    "(etc)(blvd)(cat)",
		expected: []string{
			"ebc", "eba", "ebt", "elc", "ela",
			"elt", "evc", "eva", "evt", "edc",
			"eda", "edt", "tbc", "tba", "tbt",
			"tlc", "tla", "tlt", "tvc", "tva",
			"tvt", "tdc", "tda", "tdt", "cbc",
			"cba", "cbt", "clc", "cla", "clt",
			"cvc", "cva", "cvt", "cdc", "cda",
			"cdt",
		},
	},
}
