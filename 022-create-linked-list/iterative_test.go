package createlinkedlist_test

import (
	"reflect"
	"testing"

	createlinkedlist "github.com/ju-popov/structy.net/022-create-linked-list"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := createlinkedlist.Iterative(testCase.values)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for values: '%v' is: '%v', but the actual result is: '%v'", testCase.values, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		createlinkedlist.Iterative(testCase.values)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
