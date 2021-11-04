package treesum_test

import (
	"reflect"
	"testing"

	treesum "github.com/ju-popov/structy.net/027-tree-sum"
)

func TestDepthFirstRecursive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := treesum.DepthFirstRecursive(testCase.root)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstRecursive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treesum.DepthFirstRecursive(testCase.root)
	}
}

func BenchmarkDepthFirstRecursive000(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[0]) }
func BenchmarkDepthFirstRecursive001(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[1]) }
func BenchmarkDepthFirstRecursive002(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[2]) }
