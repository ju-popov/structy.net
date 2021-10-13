package pairproduct_test

type testCase struct {
	name          string
	numbers       []int64
	targetProduct int64
	expected      [2]int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:          "test_00",
		numbers:       []int64{3, 2, 5, 4, 1},
		targetProduct: 8,
		expected:      [2]int{1, 3},
	},
	{
		name:          "test_01",
		numbers:       []int64{3, 2, 5, 4, 1},
		targetProduct: 10,
		expected:      [2]int{1, 2},
	},
	{
		name:          "test_02",
		numbers:       []int64{4, 7, 9, 2, 5, 1},
		targetProduct: 5,
		expected:      [2]int{4, 5},
	},
	{
		name:          "test_03",
		numbers:       []int64{4, 7, 9, 2, 5, 1},
		targetProduct: 35,
		expected:      [2]int{1, 4},
	},
	{
		name:          "test_04",
		numbers:       []int64{3, 2, 5, 4, 1},
		targetProduct: 10,
		expected:      [2]int{1, 2},
	},
	{
		name:          "test_05",
		numbers:       []int64{4, 6, 8, 2},
		targetProduct: 16,
		expected:      [2]int{2, 3},
	},
}
