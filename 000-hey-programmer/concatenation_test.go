package heyprogrammer_test

import (
	"reflect"
	"testing"

	heyprogrammer "github.com/ju-popov/structy.net/000-hey-programmer"
)

func TestConcatenation(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := heyprogrammer.Concatenation(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkConcatenation(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		heyprogrammer.Concatenation(tc.input)
	}
}

func BenchmarkConcatenation000(b *testing.B) { benchmarkConcatenation(b, testCases[0]) }
func BenchmarkConcatenation001(b *testing.B) { benchmarkConcatenation(b, testCases[1]) }
func BenchmarkConcatenation002(b *testing.B) { benchmarkConcatenation(b, testCases[2]) }
