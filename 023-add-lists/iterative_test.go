package addlists_test

import (
	"reflect"
	"testing"

	addlists "github.com/ju-popov/structy.net/023-add-lists"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := addlists.Iterative(tc.head1, tc.head2)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for head1: '%v' and head2: '%v' is: '%v', but the actual result is: '%v'", tc.head1, tc.head2, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		addlists.Iterative(tc.head1, tc.head2)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
func BenchmarkIterative004(b *testing.B) { benchmarkIterative(b, testCases[4]) }
