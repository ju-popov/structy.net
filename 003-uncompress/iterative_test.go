package uncompress_test

import (
	"reflect"
	"testing"

	uncompress "github.com/ju-popov/structy.net/003-uncompress"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual, err := uncompress.Iterative(tc.input)
			if err != nil {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is error: '%v'", tc.input, tc.expected, err)
			} else if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		_, err := uncompress.Iterative(tc.input)
		if err != nil {
			b.Errorf("Unexpected error: '%v'", err)
		}
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
