package uncompress_test

import (
	"reflect"
	"testing"

	uncompress "github.com/ju-popov/structy.net/003-uncompress"
)

func TestRegex(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual, err := uncompress.Regex(testCase.input)
			if err != nil {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is error: '%v'", testCase.input, testCase.expected, err)
			} else if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", testCase.input, testCase.expected, actual)
			}
		})
	}
}

func benchmarkRegex(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		_, err := uncompress.Regex(testCase.input)
		if err != nil {
			b.Errorf("Unexpected error: '%v'", err)
		}
	}
}

func BenchmarkRegex000(b *testing.B) { benchmarkRegex(b, testCases[0]) }
func BenchmarkRegex001(b *testing.B) { benchmarkRegex(b, testCases[1]) }
func BenchmarkRegex002(b *testing.B) { benchmarkRegex(b, testCases[2]) }
func BenchmarkRegex003(b *testing.B) { benchmarkRegex(b, testCases[3]) }
func BenchmarkRegex004(b *testing.B) { benchmarkRegex(b, testCases[4]) }
