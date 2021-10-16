package fivesort_test

import (
	"reflect"
	"testing"

	fivesort "github.com/ju-popov/structy.net/010-five-sort"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			numsCopy := make([]int64, len(tc.nums))
			copy(numsCopy, tc.nums)

			actual := fivesort.Simple(numsCopy)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for nums: '%v' is: '%v', but the actual result is: '%v'", tc.nums, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		fivesort.Simple(tc.nums)
	}
}

func BenchmarkSimple000(b *testing.B) { benchmarkSimple(b, testCases[0]) }
func BenchmarkSimple001(b *testing.B) { benchmarkSimple(b, testCases[1]) }
func BenchmarkSimple002(b *testing.B) { benchmarkSimple(b, testCases[2]) }
func BenchmarkSimple003(b *testing.B) { benchmarkSimple(b, testCases[3]) }
func BenchmarkSimple004(b *testing.B) { benchmarkSimple(b, testCases[4]) }
