package sumlist_test

import (
	"reflect"
	"testing"

	sumlist "github.com/ju-popov/structy.net/012-sum-list"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := sumlist.Iterative(testCase.head)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for head: '%v' is: '%v', but the actual result is: '%v'", testCase.head, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		sumlist.Iterative(testCase.head)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
