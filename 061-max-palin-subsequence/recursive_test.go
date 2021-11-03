package maxpalinsubsequence_test

import (
	"reflect"
	"testing"

	maxpalinsubsequence "github.com/ju-popov/structy.net/061-max-palin-subsequence"
)

func TestRecursive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := maxpalinsubsequence.Recursive(tc.str)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkRecursive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		maxpalinsubsequence.Recursive(tc.str)
	}
}

func BenchmarkRecursive000(b *testing.B) { benchmarkRecursive(b, testCases[0]) }
func BenchmarkRecursive001(b *testing.B) { benchmarkRecursive(b, testCases[1]) }
func BenchmarkRecursive002(b *testing.B) { benchmarkRecursive(b, testCases[2]) }
func BenchmarkRecursive003(b *testing.B) { benchmarkRecursive(b, testCases[3]) }
func BenchmarkRecursive004(b *testing.B) { benchmarkRecursive(b, testCases[4]) }
func BenchmarkRecursive005(b *testing.B) { benchmarkRecursive(b, testCases[5]) }
func BenchmarkRecursive006(b *testing.B) { benchmarkRecursive(b, testCases[6]) }
func BenchmarkRecursive007(b *testing.B) { benchmarkRecursive(b, testCases[7]) }
