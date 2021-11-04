package reverselist_test

import (
	"reflect"
	"testing"

	reverselist "github.com/ju-popov/structy.net/015-reverse-list"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := reverselist.Iterative(testCase.head.Copy())
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for head: '%v' is: '%v', but the actual result is: '%v'", testCase.head, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		reverselist.Iterative(testCase.head)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
