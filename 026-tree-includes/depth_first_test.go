package treeincludes_test

import (
	"reflect"
	"testing"

	treeincludes "github.com/ju-popov/structy.net/026-tree-includes"
)

func TestDepthFirst(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treeincludes.DepthFirst(tc.root, tc.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirst(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treeincludes.DepthFirst(tc.root, tc.target)
	}
}

func BenchmarkDepthFirst000(b *testing.B) { benchmarkDepthFirst(b, testCases[0]) }
func BenchmarkDepthFirst001(b *testing.B) { benchmarkDepthFirst(b, testCases[1]) }
func BenchmarkDepthFirst002(b *testing.B) { benchmarkDepthFirst(b, testCases[2]) }
func BenchmarkDepthFirst003(b *testing.B) { benchmarkDepthFirst(b, testCases[3]) }
func BenchmarkDepthFirst004(b *testing.B) { benchmarkDepthFirst(b, testCases[4]) }
func BenchmarkDepthFirst005(b *testing.B) { benchmarkDepthFirst(b, testCases[5]) }
