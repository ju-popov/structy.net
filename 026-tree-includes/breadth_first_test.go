package treeincludes_test

import (
	"reflect"
	"testing"

	treeincludes "github.com/ju-popov/structy.net/026-tree-includes"
)

func TestBreadthFirst(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treeincludes.BreadthFirst(tc.root, tc.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirst(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treeincludes.BreadthFirst(tc.root, tc.target)
	}
}

func BenchmarkBreadthFirst000(b *testing.B) { benchmarkBreadthFirst(b, testCases[0]) }
func BenchmarkBreadthFirst001(b *testing.B) { benchmarkBreadthFirst(b, testCases[1]) }
func BenchmarkBreadthFirst002(b *testing.B) { benchmarkBreadthFirst(b, testCases[2]) }
func BenchmarkBreadthFirst003(b *testing.B) { benchmarkBreadthFirst(b, testCases[3]) }
func BenchmarkBreadthFirst004(b *testing.B) { benchmarkBreadthFirst(b, testCases[4]) }
func BenchmarkBreadthFirst005(b *testing.B) { benchmarkBreadthFirst(b, testCases[5]) }
