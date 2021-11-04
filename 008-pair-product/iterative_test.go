package pairproduct_test

import (
	"reflect"
	"testing"

	pairproduct "github.com/ju-popov/structy.net/008-pair-product"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := pairproduct.Iterative(testCase.numbers, testCase.targetProduct)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for numbers: '%v' and targetProduct: '%v' is: '%v', but the actual result is: '%v'", testCase.numbers, testCase.targetProduct, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		pairproduct.Iterative(testCase.numbers, testCase.targetProduct)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
func BenchmarkIterative005(b *testing.B) { benchmarkIterative(b, testCases[5]) }
