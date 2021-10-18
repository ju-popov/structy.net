package treesum_test

import (
	"reflect"
	"testing"

	treesum "github.com/ju-popov/structy.net/027-tree-sum"
)

func TestDepthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treesum.DepthFirstIterative(tc.root)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treesum.DepthFirstIterative(tc.root)
	}
}

func BenchmarkDepthFirstIterative000(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[0]) }
func BenchmarkDepthFirstIterative001(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[1]) }
func BenchmarkDepthFirstIterative002(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[2]) }
