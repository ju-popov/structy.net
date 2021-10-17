package addlists_test

import (
	"reflect"
	"testing"

	addlists "github.com/ju-popov/structy.net/023-add-lists"
)

func TestRecursive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := addlists.Recursive(tc.head1, tc.head2)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for head1: '%v' and head2: '%v' is: '%v', but the actual result is: '%v'", tc.head1, tc.head2, tc.expected, actual)
			}
		})
	}
}

func benchmarkRecursive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		addlists.Recursive(tc.head1, tc.head2)
	}
}

func BenchmarkRecursive000(b *testing.B) { benchmarkRecursive(b, testCases[0]) }
func BenchmarkRecursive001(b *testing.B) { benchmarkRecursive(b, testCases[1]) }
func BenchmarkRecursive002(b *testing.B) { benchmarkRecursive(b, testCases[2]) }
func BenchmarkRecursive003(b *testing.B) { benchmarkRecursive(b, testCases[3]) }
func BenchmarkRecursive004(b *testing.B) { benchmarkRecursive(b, testCases[4]) }
