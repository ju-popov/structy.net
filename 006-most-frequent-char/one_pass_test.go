package mostfrequentchar_test

import (
	"reflect"
	"testing"

	mostfrequentchar "github.com/ju-popov/structy.net/006-most-frequent-char"
)

func TestOnePass(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := mostfrequentchar.OnePass(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkOnePass(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		mostfrequentchar.OnePass(tc.input)
	}
}

func BenchmarkOnePass000(b *testing.B) { benchmarkOnePass(b, testCases[0]) }
func BenchmarkOnePass001(b *testing.B) { benchmarkOnePass(b, testCases[1]) }
func BenchmarkOnePass002(b *testing.B) { benchmarkOnePass(b, testCases[2]) }
func BenchmarkOnePass003(b *testing.B) { benchmarkOnePass(b, testCases[3]) }
func BenchmarkOnePass004(b *testing.B) { benchmarkOnePass(b, testCases[4]) }
func BenchmarkOnePass005(b *testing.B) { benchmarkOnePass(b, testCases[5]) }
func BenchmarkOnePass006(b *testing.B) { benchmarkOnePass(b, testCases[6]) }
