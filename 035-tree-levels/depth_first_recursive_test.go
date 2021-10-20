package treelevels_test

import (
	"reflect"
	"testing"

	treelevels "github.com/ju-popov/structy.net/035-tree-levels"
)

func TestDepthFirstRecursive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treelevels.DepthFirstRecursive(tc.root)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstRecursive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treelevels.DepthFirstRecursive(tc.root)
	}
}

func BenchmarkDepthFirstRecursive000(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[0]) }
func BenchmarkDepthFirstRecursive001(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[1]) }
func BenchmarkDepthFirstRecursive002(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[2]) }
func BenchmarkDepthFirstRecursive003(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[3]) }
