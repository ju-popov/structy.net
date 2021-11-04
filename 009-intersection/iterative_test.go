package intersection_test

import (
	"reflect"
	"testing"

	intersection "github.com/ju-popov/structy.net/009-intersection"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := intersection.Iterative(testCase.a, testCase.b)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for a: '%v' and b: '%v' is: '%v', but the actual result is: '%v'", testCase.a, testCase.b, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		intersection.Iterative(testCase.a, testCase.b)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
