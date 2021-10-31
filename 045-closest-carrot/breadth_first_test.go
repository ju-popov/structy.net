package closestcarrot_test

import (
	"reflect"
	"testing"

	closestcarrot "github.com/ju-popov/structy.net/045-closest-carrot"
)

func TestBreadthFirst(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := closestcarrot.BreadthFirst(tc.grid, tc.startRow, tc.startCol)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirst(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		closestcarrot.BreadthFirst(tc.grid, tc.startRow, tc.startCol)
	}
}

func BenchmarkBreadthFirst000(b *testing.B) { benchmarkBreadthFirst(b, testCases[0]) }
func BenchmarkBreadthFirst001(b *testing.B) { benchmarkBreadthFirst(b, testCases[1]) }
func BenchmarkBreadthFirst002(b *testing.B) { benchmarkBreadthFirst(b, testCases[2]) }
func BenchmarkBreadthFirst003(b *testing.B) { benchmarkBreadthFirst(b, testCases[3]) }
func BenchmarkBreadthFirst004(b *testing.B) { benchmarkBreadthFirst(b, testCases[4]) }
