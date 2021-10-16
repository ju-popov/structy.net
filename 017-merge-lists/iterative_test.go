package mergelists_test

import (
	"reflect"
	"testing"

	mergelists "github.com/ju-popov/structy.net/017-merge-lists"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := mergelists.Iterative(tc.head1.Copy(), tc.head2.Copy())
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for head1: '%v' and head2: '%v' is: '%v', but the actual result is: '%v'", tc.head1, tc.head2, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		mergelists.Iterative(tc.head1, tc.head2)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
