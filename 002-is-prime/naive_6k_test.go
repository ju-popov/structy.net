package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

func TestNaive6K(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Naive6K(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmark6KNaive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Naive6K(tc.input)
	}
}

func BenchmarkNaive6K000(b *testing.B) { benchmark6KNaive(b, testCases[0]) }
func BenchmarkNaive6K001(b *testing.B) { benchmark6KNaive(b, testCases[1]) }
func BenchmarkNaive6K002(b *testing.B) { benchmark6KNaive(b, testCases[2]) }
func BenchmarkNaive6K003(b *testing.B) { benchmark6KNaive(b, testCases[3]) }
func BenchmarkNaive6K004(b *testing.B) { benchmark6KNaive(b, testCases[4]) }
func BenchmarkNaive6K005(b *testing.B) { benchmark6KNaive(b, testCases[5]) }
func BenchmarkNaive6K006(b *testing.B) { benchmark6KNaive(b, testCases[6]) }
func BenchmarkNaive6K007(b *testing.B) { benchmark6KNaive(b, testCases[7]) }
func BenchmarkNaive6K008(b *testing.B) { benchmark6KNaive(b, testCases[8]) }
func BenchmarkNaive6K009(b *testing.B) { benchmark6KNaive(b, testCases[9]) }
func BenchmarkNaive6K010(b *testing.B) { benchmark6KNaive(b, testCases[10]) }
func BenchmarkNaive6K011(b *testing.B) { benchmark6KNaive(b, testCases[11]) }
