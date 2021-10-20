package bottomrightvalue_test

import (
	"reflect"
	"testing"

	bottomrightvalue "github.com/ju-popov/structy.net/033-bottom-right-value"
)

func TestBreadthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := bottomrightvalue.BreadthFirstIterative(tc.root)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirstIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		bottomrightvalue.BreadthFirstIterative(tc.root)
	}
}

func BenchmarkBreadthFirstIterative000(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[0]) }
func BenchmarkBreadthFirstIterative001(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[1]) }
func BenchmarkBreadthFirstIterative002(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[2]) }
func BenchmarkBreadthFirstIterative003(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[3]) }
func BenchmarkBreadthFirstIterative004(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[4]) }
