package shortestpath_test

import (
	"reflect"
	"testing"

	shortestpath "github.com/ju-popov/structy.net/042-shortest-path"
)

func TestBreadthFirst(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := shortestpath.BreadthFirst(testCase.edges, testCase.nodeA, testCase.nodeB)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkBreadthFirst(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		shortestpath.BreadthFirst(testCase.edges, testCase.nodeA, testCase.nodeB)
	}
}

func BenchmarkBreadthFirst000(b *testing.B) { benchmarkBreadthFirst(b, testCases[0]) }
func BenchmarkBreadthFirst001(b *testing.B) { benchmarkBreadthFirst(b, testCases[1]) }
func BenchmarkBreadthFirst002(b *testing.B) { benchmarkBreadthFirst(b, testCases[2]) }
func BenchmarkBreadthFirst003(b *testing.B) { benchmarkBreadthFirst(b, testCases[3]) }
func BenchmarkBreadthFirst004(b *testing.B) { benchmarkBreadthFirst(b, testCases[4]) }
func BenchmarkBreadthFirst005(b *testing.B) { benchmarkBreadthFirst(b, testCases[5]) }
func BenchmarkBreadthFirst006(b *testing.B) { benchmarkBreadthFirst(b, testCases[6]) }
func BenchmarkBreadthFirst007(b *testing.B) { benchmarkBreadthFirst(b, testCases[7]) }
