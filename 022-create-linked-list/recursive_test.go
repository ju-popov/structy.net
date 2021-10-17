package createlinkedlist_test

import (
	"reflect"
	"testing"

	createlinkedlist "github.com/ju-popov/structy.net/022-create-linked-list"
)

func TestRecursive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := createlinkedlist.Recursive(tc.values)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for values: '%v' is: '%v', but the actual result is: '%v'", tc.values, tc.expected, actual)
			}
		})
	}
}

func benchmarkRecursive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		createlinkedlist.Recursive(tc.values)
	}
}

func BenchmarkRecursive000(b *testing.B) { benchmarkRecursive(b, testCases[0]) }
func BenchmarkRecursive001(b *testing.B) { benchmarkRecursive(b, testCases[1]) }
func BenchmarkRecursive002(b *testing.B) { benchmarkRecursive(b, testCases[2]) }
func BenchmarkRecursive003(b *testing.B) { benchmarkRecursive(b, testCases[3]) }
