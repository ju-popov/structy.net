package treeincludes_test

import (
	"reflect"
	"testing"

	treeincludes "github.com/ju-popov/structy.net/026-tree-includes"
)

func TestDepthFirstSearch(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treeincludes.DepthFirstSearch(tc.root, tc.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstSearch(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treeincludes.DepthFirstSearch(tc.root, tc.target)
	}
}

func BenchmarkDepthFirstSearch000(b *testing.B) { benchmarkDepthFirstSearch(b, testCases[0]) }
func BenchmarkDepthFirstSearch001(b *testing.B) { benchmarkDepthFirstSearch(b, testCases[1]) }
func BenchmarkDepthFirstSearch002(b *testing.B) { benchmarkDepthFirstSearch(b, testCases[2]) }
func BenchmarkDepthFirstSearch003(b *testing.B) { benchmarkDepthFirstSearch(b, testCases[3]) }
func BenchmarkDepthFirstSearch004(b *testing.B) { benchmarkDepthFirstSearch(b, testCases[4]) }
func BenchmarkDepthFirstSearch005(b *testing.B) { benchmarkDepthFirstSearch(b, testCases[5]) }
