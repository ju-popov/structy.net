package overlapsubsequence_test

import (
	"reflect"
	"testing"

	overlapsubsequence "github.com/ju-popov/structy.net/062-overlap-subsequence"
)

func TestDynamicProgramming(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := overlapsubsequence.DynamicProgramming(tc.str1, tc.str2)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkDynamicProgramming(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		overlapsubsequence.DynamicProgramming(tc.str1, tc.str2)
	}
}

func BenchmarkDynamicProgramming000(b *testing.B) { benchmarkDynamicProgramming(b, testCases[0]) }
func BenchmarkDynamicProgramming001(b *testing.B) { benchmarkDynamicProgramming(b, testCases[1]) }
func BenchmarkDynamicProgramming002(b *testing.B) { benchmarkDynamicProgramming(b, testCases[2]) }
func BenchmarkDynamicProgramming003(b *testing.B) { benchmarkDynamicProgramming(b, testCases[3]) }
func BenchmarkDynamicProgramming004(b *testing.B) { benchmarkDynamicProgramming(b, testCases[4]) }
