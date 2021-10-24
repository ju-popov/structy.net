package intersection_test

import (
	"reflect"
	"testing"

	intersection "github.com/ju-popov/structy.net/009-intersection"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := intersection.Iterative(tc.a, tc.b)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for a: '%v' and b: '%v' is: '%v', but the actual result is: '%v'", tc.a, tc.b, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		intersection.Iterative(tc.a, tc.b)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
