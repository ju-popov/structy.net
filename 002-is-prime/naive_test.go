package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

func TestNaive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Naive(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkNaive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Naive(tc.input)
	}
}

func BenchmarkNaive000(b *testing.B) { benchmarkNaive(b, testCases[0]) }
func BenchmarkNaive001(b *testing.B) { benchmarkNaive(b, testCases[1]) }
func BenchmarkNaive002(b *testing.B) { benchmarkNaive(b, testCases[2]) }
func BenchmarkNaive003(b *testing.B) { benchmarkNaive(b, testCases[3]) }
func BenchmarkNaive004(b *testing.B) { benchmarkNaive(b, testCases[4]) }
func BenchmarkNaive005(b *testing.B) { benchmarkNaive(b, testCases[5]) }
func BenchmarkNaive006(b *testing.B) { benchmarkNaive(b, testCases[6]) }
func BenchmarkNaive007(b *testing.B) { benchmarkNaive(b, testCases[7]) }
func BenchmarkNaive008(b *testing.B) { benchmarkNaive(b, testCases[8]) }
func BenchmarkNaive009(b *testing.B) { benchmarkNaive(b, testCases[9]) }
func BenchmarkNaive010(b *testing.B) { benchmarkNaive(b, testCases[10]) }
func BenchmarkNaive011(b *testing.B) { benchmarkNaive(b, testCases[11]) }
