package intersection_test

type testCase struct {
	name     string
	a        []int
	b        []int
	expected []int
}

func makeRange(min int, max int) []int {
	var result []int

	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		a:        []int{4, 2, 1, 6},
		b:        []int{3, 6, 9, 2, 10},
		expected: []int{6, 2},
	},
	{
		name:     "test_01",
		a:        []int{2, 4, 6},
		b:        []int{4, 2},
		expected: []int{4, 2},
	},
	{
		name:     "test_02",
		a:        []int{4, 2, 1},
		b:        []int{1, 2, 4, 6},
		expected: []int{1, 2, 4},
	},
	{
		name:     "test_03",
		a:        []int{0, 1, 2},
		b:        []int{10, 11},
		expected: []int{},
	},
	{
		name:     "test_04",
		a:        makeRange(0, 50000),
		b:        makeRange(0, 50000),
		expected: makeRange(0, 50000),
	},
}
