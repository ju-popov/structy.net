package undirectedpath_test

import (
	"reflect"
	"testing"

	undirectedpath "github.com/ju-popov/structy.net/039-undirected-path"
)

func TestDepthFirstRecursive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := undirectedpath.DepthFirstRecursive(tc.edges, tc.nodeA, tc.nodeB)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstRecursive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		undirectedpath.DepthFirstRecursive(tc.edges, tc.nodeA, tc.nodeB)
	}
}

func BenchmarkDepthFirstRecursive000(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[0]) }
func BenchmarkDepthFirstRecursive001(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[1]) }
func BenchmarkDepthFirstRecursive002(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[2]) }
func BenchmarkDepthFirstRecursive003(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[3]) }
func BenchmarkDepthFirstRecursive004(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[4]) }
func BenchmarkDepthFirstRecursive005(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[5]) }
func BenchmarkDepthFirstRecursive006(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[6]) }
func BenchmarkDepthFirstRecursive007(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[7]) }
func BenchmarkDepthFirstRecursive008(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[8]) }
