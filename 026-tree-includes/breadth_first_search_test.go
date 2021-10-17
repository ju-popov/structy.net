package treeincludes_test

import (
	"reflect"
	"testing"

	treeincludes "github.com/ju-popov/structy.net/026-tree-includes"
)

func TestBreadthFirstSearch(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treeincludes.BreadthFirstSearch(tc.root, tc.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirstSearch(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treeincludes.BreadthFirstSearch(tc.root, tc.target)
	}
}

func BenchmarkBreadthFirstSearch000(b *testing.B) { benchmarkBreadthFirstSearch(b, testCases[0]) }
func BenchmarkBreadthFirstSearch001(b *testing.B) { benchmarkBreadthFirstSearch(b, testCases[1]) }
func BenchmarkBreadthFirstSearch002(b *testing.B) { benchmarkBreadthFirstSearch(b, testCases[2]) }
func BenchmarkBreadthFirstSearch003(b *testing.B) { benchmarkBreadthFirstSearch(b, testCases[3]) }
func BenchmarkBreadthFirstSearch004(b *testing.B) { benchmarkBreadthFirstSearch(b, testCases[4]) }
func BenchmarkBreadthFirstSearch005(b *testing.B) { benchmarkBreadthFirstSearch(b, testCases[5]) }
