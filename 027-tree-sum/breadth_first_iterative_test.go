package treesum_test

import (
	"reflect"
	"testing"

	treesum "github.com/ju-popov/structy.net/027-tree-sum"
)

func TestBreadthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := treesum.BreadthFirstIterative(testCase.root)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirstIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treesum.BreadthFirstIterative(testCase.root)
	}
}

func BenchmarkBreadthFirstIterative000(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[0]) }
func BenchmarkBreadthFirstIterative001(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[1]) }
func BenchmarkBreadthFirstIterative002(b *testing.B) { benchmarkBreadthFirstIterative(b, testCases[2]) }
