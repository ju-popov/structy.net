package getnodevalue_test

import (
	"reflect"
	"testing"

	getnodevalue "github.com/ju-popov/structy.net/014-get-node-value"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := getnodevalue.Simple(tc.head, tc.index)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for test '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		getnodevalue.Simple(tc.head, tc.index)
	}
}

func BenchmarkSimple000(b *testing.B) { benchmarkSimple(b, testCases[0]) }
func BenchmarkSimple001(b *testing.B) { benchmarkSimple(b, testCases[1]) }
func BenchmarkSimple002(b *testing.B) { benchmarkSimple(b, testCases[2]) }
func BenchmarkSimple003(b *testing.B) { benchmarkSimple(b, testCases[3]) }
func BenchmarkSimple004(b *testing.B) { benchmarkSimple(b, testCases[4]) }
