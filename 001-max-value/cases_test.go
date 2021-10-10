package maxvalue_test

type testCase struct {
	name     string
	input    []float64
	expected float64
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		input:    []float64{4, 7, 2, 8, 10, 9},
		expected: 10,
	},
	{
		name:     "test_01",
		input:    []float64{10, 5, 40, 40.3},
		expected: 40.3,
	},
	{
		name:     "test_02",
		input:    []float64{-5, -2, -1, -11},
		expected: -1,
	},
	{
		name:     "test_03",
		input:    []float64{42},
		expected: 42,
	},
	{
		name:     "test_04",
		input:    []float64{1000, 8},
		expected: 1000,
	},
	{
		name:     "test_05",
		input:    []float64{1000, 8, 9000},
		expected: 9000,
	},
	{
		name:     "test_06",
		input:    []float64{2, 5, 1, 1, 4},
		expected: 5,
	},
}
