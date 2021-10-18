package treevaluecount_test

import (
	"reflect"
	"testing"

	treevaluecount "github.com/ju-popov/structy.net/031-tree-value-count"
)

func TestDepthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treevaluecount.DepthFirstIterative(tc.root, tc.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treevaluecount.DepthFirstIterative(tc.root, tc.target)
	}
}

func BenchmarkDepthFirstIterative000(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[0]) }
func BenchmarkDepthFirstIterative001(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[1]) }
func BenchmarkDepthFirstIterative002(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[2]) }
func BenchmarkDepthFirstIterative003(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[3]) }
func BenchmarkDepthFirstIterative004(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[4]) }
