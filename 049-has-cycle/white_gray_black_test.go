package hascycle_test

import (
	"reflect"
	"testing"

	hascycle "github.com/ju-popov/structy.net/049-has-cycle"
)

func TestWhiteGreyBlack(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := hascycle.WhiteGreyBlack(tc.graph)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkWhiteGreyBlack(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		hascycle.WhiteGreyBlack(tc.graph)
	}
}

func BenchmarkWhiteGreyBlack000(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[0]) }
func BenchmarkWhiteGreyBlack001(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[1]) }
func BenchmarkWhiteGreyBlack002(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[2]) }
func BenchmarkWhiteGreyBlack003(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[3]) }
func BenchmarkWhiteGreyBlack004(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[4]) }
