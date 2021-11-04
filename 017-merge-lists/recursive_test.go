package mergelists_test

import (
	"reflect"
	"testing"

	mergelists "github.com/ju-popov/structy.net/017-merge-lists"
)

func TestRecursive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := mergelists.Recursive(testCase.head1.Copy(), testCase.head2.Copy())
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkRecursive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		mergelists.Recursive(testCase.head1.Copy(), testCase.head2.Copy())
	}
}

func BenchmarkRecursive000(b *testing.B) { benchmarkRecursive(b, testCases[0]) }
func BenchmarkRecursive001(b *testing.B) { benchmarkRecursive(b, testCases[1]) }
func BenchmarkRecursive002(b *testing.B) { benchmarkRecursive(b, testCases[2]) }
