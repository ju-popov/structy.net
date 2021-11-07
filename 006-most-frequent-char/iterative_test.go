package mostfrequentchar_test

import (
	"reflect"
	"testing"

	mostfrequentchar "github.com/ju-popov/structy.net/006-most-frequent-char"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := mostfrequentchar.Iterative(testCase.s)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		mostfrequentchar.Iterative(testCase.s)
	}
}

func BenchmarkIterative000(b *testing.B)  { benchmarkIterative(b, testCases[0]) }
func BenchmarkIteratives001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B)  { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B)  { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B)  { benchmarkIterative(b, testCases[4]) }
func BenchmarkIterative005(b *testing.B)  { benchmarkIterative(b, testCases[5]) }
func BenchmarkIterative006(b *testing.B)  { benchmarkIterative(b, testCases[6]) }
