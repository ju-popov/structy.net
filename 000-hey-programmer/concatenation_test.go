package heyprogrammer_test

import (
	"reflect"
	"testing"

	heyprogrammer "github.com/ju-popov/structy.net/000-hey-programmer"
)

func TestConcatenation(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := heyprogrammer.Concatenation(testCase.input)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", testCase.input, testCase.expected, actual)
			}
		})
	}
}

func benchmarkConcatenation(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		heyprogrammer.Concatenation(testCase.input)
	}
}

func BenchmarkConcatenation000(b *testing.B) { benchmarkConcatenation(b, testCases[0]) }
func BenchmarkConcatenation001(b *testing.B) { benchmarkConcatenation(b, testCases[1]) }
func BenchmarkConcatenation002(b *testing.B) { benchmarkConcatenation(b, testCases[2]) }
