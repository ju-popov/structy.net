package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

func TestNaive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Naive(testCase.n)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkNaive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Naive(testCase.n)
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
