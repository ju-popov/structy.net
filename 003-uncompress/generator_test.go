package uncompress_test

import (
	"reflect"
	"testing"

	uncompress "github.com/ju-popov/structy.net/003-uncompress"
)

func TestGenerator(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := uncompress.Generator(testCase.input)
			if err != nil {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is error: '%v'", testCase.input, testCase.expected, err)
			} else if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is: '%v'", testCase.input, testCase.expected, actual)
			}
		})
	}
}

func benchmarkGenerator(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		_, err := uncompress.Generator(testCase.input)
		if err != nil {
			b.Errorf("Unexpected error: '%v'", err)
		}
	}
}

func BenchmarkGenerator000(b *testing.B) { benchmarkGenerator(b, testCases[0]) }
func BenchmarkGenerator001(b *testing.B) { benchmarkGenerator(b, testCases[1]) }
func BenchmarkGenerator002(b *testing.B) { benchmarkGenerator(b, testCases[2]) }
func BenchmarkGenerator003(b *testing.B) { benchmarkGenerator(b, testCases[3]) }
func BenchmarkGenerator004(b *testing.B) { benchmarkGenerator(b, testCases[4]) }
