package anagrams_test

import (
	"reflect"
	"testing"

	anagrams "github.com/ju-popov/structy.net/005-anagrams"
)

func TestSimple(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := anagrams.Simple(tc.s1, tc.s2)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input: '%v', '%v' is: '%v', but the actual result is: '%v'", tc.s1, tc.s2, tc.expected, actual)
			}
		})
	}
}

func benchmarkSimple(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		anagrams.Simple(tc.s1, tc.s2)
	}
}

func BenchmarkSimple000(b *testing.B) { benchmarkSimple(b, testCases[0]) }
func BenchmarkSimple001(b *testing.B) { benchmarkSimple(b, testCases[1]) }
func BenchmarkSimple002(b *testing.B) { benchmarkSimple(b, testCases[2]) }
func BenchmarkSimple003(b *testing.B) { benchmarkSimple(b, testCases[3]) }
func BenchmarkSimple004(b *testing.B) { benchmarkSimple(b, testCases[4]) }
func BenchmarkSimple005(b *testing.B) { benchmarkSimple(b, testCases[5]) }
func BenchmarkSimple006(b *testing.B) { benchmarkSimple(b, testCases[6]) }
func BenchmarkSimple007(b *testing.B) { benchmarkSimple(b, testCases[7]) }
func BenchmarkSimple008(b *testing.B) { benchmarkSimple(b, testCases[8]) }
