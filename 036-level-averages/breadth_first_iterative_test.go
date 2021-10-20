package levelaverages_test

import (
	"reflect"
	"testing"

	levelaverages "github.com/ju-popov/structy.net/036-level-averages"
)

func TestBreadthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := levelaverages.BreadthFirstIterative(tc.root)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirstIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		levelaverages.BreadthFirstIterative(tc.root)
	}
}

func BenchmarkBreadthFirstIterative000(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[0]) }
func BenchmarkBreadthFirstIterative001(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[1]) }
func BenchmarkBreadthFirstIterative002(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[2]) }
func BenchmarkBreadthFirstIterative003(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[3]) }
func BenchmarkBreadthFirstIterative004(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[4]) }