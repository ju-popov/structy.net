package maxpalinsubsequence_test

import (
	"reflect"
	"testing"

	maxpalinsubsequence "github.com/ju-popov/structy.net/061-max-palin-subsequence"
)

func TestDynamicProgramming(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := maxpalinsubsequence.DynamicProgramming(testCase.str)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDynamicProgramming(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		maxpalinsubsequence.DynamicProgramming(testCase.str)
	}
}

func BenchmarkDynamicProgramming000(b *testing.B) { benchmarkDynamicProgramming(b, testCases[0]) }
func BenchmarkDynamicProgramming001(b *testing.B) { benchmarkDynamicProgramming(b, testCases[1]) }
func BenchmarkDynamicProgramming002(b *testing.B) { benchmarkDynamicProgramming(b, testCases[2]) }
func BenchmarkDynamicProgramming003(b *testing.B) { benchmarkDynamicProgramming(b, testCases[3]) }
func BenchmarkDynamicProgramming004(b *testing.B) { benchmarkDynamicProgramming(b, testCases[4]) }
func BenchmarkDynamicProgramming005(b *testing.B) { benchmarkDynamicProgramming(b, testCases[5]) }
func BenchmarkDynamicProgramming006(b *testing.B) { benchmarkDynamicProgramming(b, testCases[6]) }
func BenchmarkDynamicProgramming007(b *testing.B) { benchmarkDynamicProgramming(b, testCases[7]) }
