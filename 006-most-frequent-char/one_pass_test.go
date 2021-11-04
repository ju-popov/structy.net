package mostfrequentchar_test

import (
	"reflect"
	"testing"

	mostfrequentchar "github.com/ju-popov/structy.net/006-most-frequent-char"
)

func TestOnePass(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := mostfrequentchar.OnePass(testCase.input)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", testCase.input, testCase.expected, actual)
			}
		})
	}
}

func benchmarkOnePass(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		mostfrequentchar.OnePass(testCase.input)
	}
}

func BenchmarkOnePass000(b *testing.B) { benchmarkOnePass(b, testCases[0]) }
func BenchmarkOnePass001(b *testing.B) { benchmarkOnePass(b, testCases[1]) }
func BenchmarkOnePass002(b *testing.B) { benchmarkOnePass(b, testCases[2]) }
func BenchmarkOnePass003(b *testing.B) { benchmarkOnePass(b, testCases[3]) }
func BenchmarkOnePass004(b *testing.B) { benchmarkOnePass(b, testCases[4]) }
func BenchmarkOnePass005(b *testing.B) { benchmarkOnePass(b, testCases[5]) }
func BenchmarkOnePass006(b *testing.B) { benchmarkOnePass(b, testCases[6]) }
