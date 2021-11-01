package semestersrequired_test

type testCase struct {
	name       string
	numCourses int
	prereqs    [][2]int
	expected   int
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:       "test_00",
		numCourses: 6,
		prereqs: [][2]int{
			{1, 2},
			{2, 4},
			{3, 5},
			{0, 5},
		},
		expected: 3,
	},
	{
		name:       "test_01",
		numCourses: 7,
		prereqs: [][2]int{
			{4, 3},
			{3, 2},
			{2, 1},
			{1, 0},
			{5, 2},
			{5, 6},
		},
		expected: 5,
	},
	{
		name:       "test_02",
		numCourses: 5,
		prereqs: [][2]int{
			{1, 0},
			{3, 4},
			{1, 2},
			{3, 2},
		},
		expected: 2,
	},
	{
		name:       "test_03",
		numCourses: 12,
		prereqs:    [][2]int{},
		expected:   1,
	},
	{
		name:       "test_04",
		numCourses: 3,
		prereqs: [][2]int{
			{0, 2},
			{0, 1},
			{1, 2},
		},
		expected: 3,
	},
	{
		name:       "test_05",
		numCourses: 6,
		prereqs: [][2]int{
			{3, 4},
			{3, 0},
			{3, 1},
			{3, 2},
			{3, 5},
		},
		expected: 2,
	},
}
