package semestersrequired_test

import (
	"reflect"
	"testing"

	semestersrequired "github.com/ju-popov/structy.net/047-semesters-required"
)

func TestDepthFirstRecursive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := semestersrequired.DepthFirstRecursive(testCase.numCourses, testCase.prereqs)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstRecursive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		semestersrequired.DepthFirstRecursive(testCase.numCourses, testCase.prereqs)
	}
}

func BenchmarkDepthFirstRecursive000(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[0]) }
func BenchmarkDepthFirstRecursive001(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[1]) }
func BenchmarkDepthFirstRecursive002(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[2]) }
func BenchmarkDepthFirstRecursive003(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[3]) }
func BenchmarkDepthFirstRecursive004(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[4]) }
func BenchmarkDepthFirstRecursive005(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[5]) }
