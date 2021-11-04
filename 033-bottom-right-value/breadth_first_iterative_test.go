package bottomrightvalue_test

import (
	"reflect"
	"testing"

	bottomrightvalue "github.com/ju-popov/structy.net/033-bottom-right-value"
)

func TestBreadthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := bottomrightvalue.BreadthFirstIterative(testCase.root)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirstIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		bottomrightvalue.BreadthFirstIterative(testCase.root)
	}
}

func BenchmarkBreadthFirstIterative000(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[0]) }
func BenchmarkBreadthFirstIterative001(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[1]) }
func BenchmarkBreadthFirstIterative002(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[2]) }
func BenchmarkBreadthFirstIterative003(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[3]) }
func BenchmarkBreadthFirstIterative004(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[4]) }
