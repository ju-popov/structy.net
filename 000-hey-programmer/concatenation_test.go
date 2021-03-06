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
			actual := heyprogrammer.Concatenation(testCase.s)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkConcatenation(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		heyprogrammer.Concatenation(testCase.s)
	}
}

func BenchmarkConcatenation000(b *testing.B) { benchmarkConcatenation(b, testCases[0]) }
func BenchmarkConcatenation001(b *testing.B) { benchmarkConcatenation(b, testCases[1]) }
func BenchmarkConcatenation002(b *testing.B) { benchmarkConcatenation(b, testCases[2]) }
