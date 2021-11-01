package prereqspossible_test

type testCase struct {
	name       string
	numCourses int
	prereqs    [][2]int
	expected   bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:       "test_00",
		numCourses: 6,
		prereqs: [][2]int{
			{0, 1},
			{2, 3},
			{0, 2},
			{1, 3},
			{4, 5},
		},
		expected: true,
	},
	{
		name:       "test_01",
		numCourses: 6,
		prereqs: [][2]int{
			{0, 1},
			{2, 3},
			{0, 2},
			{1, 3},
			{4, 5},
			{3, 0},
		},
		expected: false,
	},
	{
		name:       "test_02",
		numCourses: 5,
		prereqs: [][2]int{
			{2, 4},
			{1, 0},
			{0, 2},
			{0, 4},
		},
		expected: true,
	},
	{
		name:       "test_03",
		numCourses: 6,
		prereqs: [][2]int{
			{2, 4},
			{1, 0},
			{0, 2},
			{0, 4},
			{5, 3},
			{3, 5},
		},
		expected: false,
	},
	{
		name:       "test_04",
		numCourses: 8,
		prereqs: [][2]int{
			{1, 0},
			{0, 6},
			{2, 0},
			{0, 5},
			{3, 7},
			{4, 3},
		},
		expected: true,
	},
	{
		name:       "test_05",
		numCourses: 8,
		prereqs: [][2]int{
			{1, 0},
			{0, 6},
			{2, 0},
			{0, 5},
			{3, 7},
			{7, 4},
			{4, 3},
		},
		expected: false,
	},
	{
		name:       "test_06",
		numCourses: 42,
		prereqs: [][2]int{
			{6, 36},
		},
		expected: true,
	},
}
