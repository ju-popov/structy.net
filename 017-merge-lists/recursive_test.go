package mergelists_test

import (
	"reflect"
	"testing"

	mergelists "github.com/ju-popov/structy.net/017-merge-lists"
)

func TestRecursive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := mergelists.Recursive(tc.head1.Copy(), tc.head2.Copy())
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for head1: '%v' and head2: '%v' is: '%v', but the actual result is: '%v'", tc.head1, tc.head2, tc.expected, actual)
			}
		})
	}
}

func benchmarkRecursive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		mergelists.Recursive(tc.head1, tc.head2)
	}
}

func BenchmarkRecursive000(b *testing.B) { benchmarkRecursive(b, testCases[0]) }
func BenchmarkRecursive001(b *testing.B) { benchmarkRecursive(b, testCases[1]) }
func BenchmarkRecursive002(b *testing.B) { benchmarkRecursive(b, testCases[2]) }
