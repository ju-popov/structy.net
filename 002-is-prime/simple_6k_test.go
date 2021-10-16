package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

func TestSimple6K(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Simple6K(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple6K(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Simple6K(tc.input)
	}
}

func BenchmarkSimple6K000(b *testing.B) { benchmarkSimple6K(b, testCases[0]) }
func BenchmarkSimple6K001(b *testing.B) { benchmarkSimple6K(b, testCases[1]) }
func BenchmarkSimple6K002(b *testing.B) { benchmarkSimple6K(b, testCases[2]) }
func BenchmarkSimple6K003(b *testing.B) { benchmarkSimple6K(b, testCases[3]) }
func BenchmarkSimple6K004(b *testing.B) { benchmarkSimple6K(b, testCases[4]) }
func BenchmarkSimple6K005(b *testing.B) { benchmarkSimple6K(b, testCases[5]) }
func BenchmarkSimple6K006(b *testing.B) { benchmarkSimple6K(b, testCases[6]) }
func BenchmarkSimple6K007(b *testing.B) { benchmarkSimple6K(b, testCases[7]) }
func BenchmarkSimple6K008(b *testing.B) { benchmarkSimple6K(b, testCases[8]) }
func BenchmarkSimple6K009(b *testing.B) { benchmarkSimple6K(b, testCases[9]) }
func BenchmarkSimple6K010(b *testing.B) { benchmarkSimple6K(b, testCases[10]) }
func BenchmarkSimple6K011(b *testing.B) { benchmarkSimple6K(b, testCases[11]) }
