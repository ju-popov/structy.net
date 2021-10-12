package intersection_test

type testCase struct {
	name     string
	a        []int64
	b        []int64
	expected []int64
}

func makeRange(min int64, max int64) []int64 {
	var result []int64

	for i := min; i < max; i++ {
		result = append(result, i)
	}

	return result
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		a:        []int64{4, 2, 1, 6},
		b:        []int64{3, 6, 9, 2, 10},
		expected: []int64{6, 2},
	},
	{
		name:     "test_01",
		a:        []int64{2, 4, 6},
		b:        []int64{4, 2},
		expected: []int64{4, 2},
	},
	{
		name:     "test_02",
		a:        []int64{4, 2, 1},
		b:        []int64{1, 2, 4, 6},
		expected: []int64{1, 2, 4},
	},
	{
		name:     "test_03",
		a:        []int64{0, 1, 2},
		b:        []int64{10, 11},
		expected: []int64{},
	},
	{
		name:     "test_04",
		a:        makeRange(0, 50000),
		b:        makeRange(0, 50000),
		expected: makeRange(0, 50000),
	},
}
