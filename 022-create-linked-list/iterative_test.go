package createlinkedlist_test

import (
	"reflect"
	"testing"

	createlinkedlist "github.com/ju-popov/structy.net/022-create-linked-list"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := createlinkedlist.Iterative(tc.values)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for values: '%v' is: '%v', but the actual result is: '%v'", tc.values, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		createlinkedlist.Iterative(tc.values)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
