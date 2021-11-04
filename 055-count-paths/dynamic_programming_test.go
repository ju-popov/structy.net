package countpaths_test

import (
	"reflect"
	"testing"

	countpaths "github.com/ju-popov/structy.net/055-count-paths"
)

func TestDynamicProgramming(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := countpaths.DynamicProgramming(testCase.grid)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDynamicProgramming(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		countpaths.DynamicProgramming(testCase.grid)
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
func BenchmarkDynamicProgramming008(b *testing.B) { benchmarkDynamicProgramming(b, testCases[8]) }
