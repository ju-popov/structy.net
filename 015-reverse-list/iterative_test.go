package reverselist_test

import (
	"reflect"
	"testing"

	reverselist "github.com/ju-popov/structy.net/015-reverse-list"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := reverselist.Iterative(tc.head.Copy())
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for head: '%v' is: '%v', but the actual result is: '%v'", tc.head, tc.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		reverselist.Iterative(tc.head)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
