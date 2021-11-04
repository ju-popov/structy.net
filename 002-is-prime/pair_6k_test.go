package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

func TestPair6K(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Pair6K(testCase.input)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkPair6K(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Pair6K(testCase.input)
	}
}

func BenchmarkPair6K000(b *testing.B) { benchmarkPair6K(b, testCases[0]) }
func BenchmarkPair6K001(b *testing.B) { benchmarkPair6K(b, testCases[1]) }
func BenchmarkPair6K002(b *testing.B) { benchmarkPair6K(b, testCases[2]) }
func BenchmarkPair6K003(b *testing.B) { benchmarkPair6K(b, testCases[3]) }
func BenchmarkPair6K004(b *testing.B) { benchmarkPair6K(b, testCases[4]) }
func BenchmarkPair6K005(b *testing.B) { benchmarkPair6K(b, testCases[5]) }
func BenchmarkPair6K006(b *testing.B) { benchmarkPair6K(b, testCases[6]) }
func BenchmarkPair6K007(b *testing.B) { benchmarkPair6K(b, testCases[7]) }
func BenchmarkPair6K008(b *testing.B) { benchmarkPair6K(b, testCases[8]) }
func BenchmarkPair6K009(b *testing.B) { benchmarkPair6K(b, testCases[9]) }
func BenchmarkPair6K010(b *testing.B) { benchmarkPair6K(b, testCases[10]) }
func BenchmarkPair6K011(b *testing.B) { benchmarkPair6K(b, testCases[11]) }
