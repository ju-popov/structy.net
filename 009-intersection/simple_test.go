package intersection_test

import (
	"reflect"
	"testing"

	intersection "github.com/ju-popov/structy.net/009-intersection"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := intersection.Simple(tc.a, tc.b)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for a '%v' and b '%v' is: '%v', but the actual result is: '%v'", tc.a, tc.b, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		intersection.Simple(tc.a, tc.b)
	}
}

func BenchmarkSimple000(b *testing.B) { benchmarkSimple(b, testCases[0]) }
func BenchmarkSimple001(b *testing.B) { benchmarkSimple(b, testCases[1]) }
func BenchmarkSimple002(b *testing.B) { benchmarkSimple(b, testCases[2]) }
func BenchmarkSimple003(b *testing.B) { benchmarkSimple(b, testCases[3]) }
func BenchmarkSimple004(b *testing.B) { benchmarkSimple(b, testCases[4]) }
