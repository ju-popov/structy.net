package compress_test

import (
	"reflect"
	"testing"

	compress "github.com/ju-popov/structy.net/004-compress"
)

func TestGenerator(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := compress.Generator(testCase.input)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkGenerator(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		compress.Generator(testCase.input)
	}
}

func BenchmarkGenerator000(b *testing.B) { benchmarkGenerator(b, testCases[0]) }
func BenchmarkGenerator001(b *testing.B) { benchmarkGenerator(b, testCases[1]) }
func BenchmarkGenerator002(b *testing.B) { benchmarkGenerator(b, testCases[2]) }
func BenchmarkGenerator003(b *testing.B) { benchmarkGenerator(b, testCases[3]) }
func BenchmarkGenerator004(b *testing.B) { benchmarkGenerator(b, testCases[4]) }
