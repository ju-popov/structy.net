package treeminvalue_test

import (
	"reflect"
	"testing"

	treeminvalue "github.com/ju-popov/structy.net/028-tree-min-value"
)

func TestBreadthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := treeminvalue.BreadthFirstIterative(tc.root)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirstIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treeminvalue.BreadthFirstIterative(tc.root)
	}
}

func BenchmarkBreadthFirstIterative000(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[0]) }
func BenchmarkBreadthFirstIterative001(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[1]) }
func BenchmarkBreadthFirstIterative002(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[2]) }
