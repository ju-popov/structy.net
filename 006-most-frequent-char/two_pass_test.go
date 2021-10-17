package mostfrequentchar_test

import (
	"reflect"
	"testing"

	mostfrequentchar "github.com/ju-popov/structy.net/006-most-frequent-char"
)

func TestTwoPass(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := mostfrequentchar.TwoPass(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkTwoPass(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		mostfrequentchar.TwoPass(tc.input)
	}
}

func BenchmarkTwoPass000(b *testing.B) { benchmarkTwoPass(b, testCases[0]) }
func BenchmarkTwoPass001(b *testing.B) { benchmarkTwoPass(b, testCases[1]) }
func BenchmarkTwoPass002(b *testing.B) { benchmarkTwoPass(b, testCases[2]) }
func BenchmarkTwoPass003(b *testing.B) { benchmarkTwoPass(b, testCases[3]) }
func BenchmarkTwoPass004(b *testing.B) { benchmarkTwoPass(b, testCases[4]) }
func BenchmarkTwoPass005(b *testing.B) { benchmarkTwoPass(b, testCases[5]) }
func BenchmarkTwoPass006(b *testing.B) { benchmarkTwoPass(b, testCases[6]) }