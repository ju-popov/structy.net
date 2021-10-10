package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Simple(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Simple(tc.input)
	}
}

func BenchmarkSimple000(b *testing.B) { benchmarkSimple(b, testCases[0]) }
func BenchmarkSimple001(b *testing.B) { benchmarkSimple(b, testCases[1]) }
func BenchmarkSimple002(b *testing.B) { benchmarkSimple(b, testCases[2]) }
func BenchmarkSimple003(b *testing.B) { benchmarkSimple(b, testCases[3]) }
func BenchmarkSimple004(b *testing.B) { benchmarkSimple(b, testCases[4]) }
func BenchmarkSimple005(b *testing.B) { benchmarkSimple(b, testCases[5]) }
func BenchmarkSimple006(b *testing.B) { benchmarkSimple(b, testCases[6]) }
func BenchmarkSimple007(b *testing.B) { benchmarkSimple(b, testCases[7]) }
func BenchmarkSimple008(b *testing.B) { benchmarkSimple(b, testCases[8]) }
func BenchmarkSimple009(b *testing.B) { benchmarkSimple(b, testCases[9]) }
func BenchmarkSimple010(b *testing.B) { benchmarkSimple(b, testCases[10]) }
func BenchmarkSimple011(b *testing.B) { benchmarkSimple(b, testCases[11]) }
