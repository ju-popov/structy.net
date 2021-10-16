package pairproduct_test

import (
	"reflect"
	"testing"

	pairproduct "github.com/ju-popov/structy.net/008-pair-product"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := pairproduct.Simple(tc.numbers, tc.targetProduct)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for numbers: '%v' and targetProduct: '%v' is: '%v', but the actual result is: '%v'", tc.numbers, tc.targetProduct, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		pairproduct.Simple(tc.numbers, tc.targetProduct)
	}
}

func BenchmarkSimple000(b *testing.B) { benchmarkSimple(b, testCases[0]) }
func BenchmarkSimple001(b *testing.B) { benchmarkSimple(b, testCases[1]) }
func BenchmarkSimple002(b *testing.B) { benchmarkSimple(b, testCases[2]) }
func BenchmarkSimple003(b *testing.B) { benchmarkSimple(b, testCases[3]) }
func BenchmarkSimple004(b *testing.B) { benchmarkSimple(b, testCases[4]) }
func BenchmarkSimple005(b *testing.B) { benchmarkSimple(b, testCases[5]) }
