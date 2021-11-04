package tribonacci_test

import (
	"reflect"
	"testing"

	tribonacci "github.com/ju-popov/structy.net/052-tribonacci"
)

func TestRecursive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := tribonacci.Recursive(testCase.n)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkRecursive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		tribonacci.Recursive(testCase.n)
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
