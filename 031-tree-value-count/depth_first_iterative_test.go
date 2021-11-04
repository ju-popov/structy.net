package treevaluecount_test

import (
	"reflect"
	"testing"

	treevaluecount "github.com/ju-popov/structy.net/031-tree-value-count"
)

func TestDepthFirstIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := treevaluecount.DepthFirstIterative(testCase.root, testCase.target)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		treevaluecount.DepthFirstIterative(testCase.root, testCase.target)
	}
}

func BenchmarkDepthFirstIterative000(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[0]) }
func BenchmarkDepthFirstIterative001(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[1]) }
func BenchmarkDepthFirstIterative002(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[2]) }
func BenchmarkDepthFirstIterative003(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[3]) }
func BenchmarkDepthFirstIterative004(b *testing.B) { benchmarkDepthFirstIterative(b, testCases[4]) }
