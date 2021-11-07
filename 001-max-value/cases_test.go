package maxvalue_test

type testCase struct {
	name     string
	nums     []float64
	expected float64
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		nums:     []float64{4, 7, 2, 8, 10, 9},
		expected: 10,
	},
	{
		name:     "test_01",
		nums:     []float64{10, 5, 40, 40.3},
		expected: 40.3,
	},
	{
		name:     "test_02",
		nums:     []float64{-5, -2, -1, -11},
		expected: -1,
	},
	{
		name:     "test_03",
		nums:     []float64{42},
		expected: 42,
	},
	{
		name:     "test_04",
		nums:     []float64{1000, 8},
		expected: 1000,
	},
	{
		name:     "test_05",
		nums:     []float64{1000, 8, 9000},
		expected: 9000,
	},
	{
		name:     "test_06",
		nums:     []float64{2, 5, 1, 1, 4},
		expected: 5,
	},
}
