package treeminvalue_test

import (
	"reflect"
	"testing"

	treeminvalue "github.com/ju-popov/structy.net/028-tree-min-value"
)

func TestDepthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treeminvalue.DepthFirstIterative(tc.root)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treeminvalue.DepthFirstIterative(tc.root)
	}
}

func BenchmarkDepthFirstIterative000(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[0]) }
func BenchmarkDepthFirstIterative001(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[1]) }
func BenchmarkDepthFirstIterative002(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[2]) }