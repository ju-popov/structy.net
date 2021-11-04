package getnodevalue_test

import (
	"reflect"
	"testing"

	getnodevalue "github.com/ju-popov/structy.net/014-get-node-value"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := getnodevalue.Iterative(testCase.head, testCase.index)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for head: '%v' and index: '%v' is: '%v', but the actual result is: '%v'", testCase.head, testCase.index, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		getnodevalue.Iterative(testCase.head, testCase.index)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
