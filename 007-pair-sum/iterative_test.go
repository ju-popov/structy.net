package pairsum_test

import (
	"reflect"
	"testing"

	pairsum "github.com/ju-popov/structy.net/007-pair-sum"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := pairsum.Iterative(tc.numbers, tc.targetSum)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for numbers: '%v' and targetSum: '%v' is: '%v', but the actual result is: '%v'", tc.numbers, tc.targetSum, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		pairsum.Iterative(tc.numbers, tc.targetSum)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
func BenchmarkIterative005(b *testing.B) { benchmarkIterative(b, testCases[5]) }
