package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

func TestPair(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Pair(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkPair(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Pair(tc.input)
	}
}

func BenchmarkPair000(b *testing.B) { benchmarkPair(b, testCases[0]) }
func BenchmarkPair001(b *testing.B) { benchmarkPair(b, testCases[1]) }
func BenchmarkPair002(b *testing.B) { benchmarkPair(b, testCases[2]) }
func BenchmarkPair003(b *testing.B) { benchmarkPair(b, testCases[3]) }
func BenchmarkPair004(b *testing.B) { benchmarkPair(b, testCases[4]) }
func BenchmarkPair005(b *testing.B) { benchmarkPair(b, testCases[5]) }
func BenchmarkPair006(b *testing.B) { benchmarkPair(b, testCases[6]) }
func BenchmarkPair007(b *testing.B) { benchmarkPair(b, testCases[7]) }
func BenchmarkPair008(b *testing.B) { benchmarkPair(b, testCases[8]) }
func BenchmarkPair009(b *testing.B) { benchmarkPair(b, testCases[9]) }
func BenchmarkPair010(b *testing.B) { benchmarkPair(b, testCases[10]) }
func BenchmarkPair011(b *testing.B) { benchmarkPair(b, testCases[11]) }