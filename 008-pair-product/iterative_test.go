package pairproduct_test

import (
	"reflect"
	"testing"

	pairproduct "github.com/ju-popov/structy.net/008-pair-product"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := pairproduct.Iterative(tc.numbers, tc.targetProduct)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for numbers: '%v' and targetProduct: '%v' is: '%v', but the actual result is: '%v'", tc.numbers, tc.targetProduct, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		pairproduct.Iterative(tc.numbers, tc.targetProduct)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
func BenchmarkIterative005(b *testing.B) { benchmarkIterative(b, testCases[5]) }
