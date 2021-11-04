package insertnode_test

import (
	"reflect"
	"testing"

	insertnode "github.com/ju-popov/structy.net/021-insert-node"
)

func TestIterative(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := insertnode.Iterative(testCase.head.Copy(), testCase.value, testCase.index)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for head: '%v' value: '%v' and index: '%v' is: '%v', but the actual result is: '%v'", testCase.head, testCase.value, testCase.index, testCase.expected, actual)
			}
		})
	}
}

func benchmarkIterative(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		insertnode.Iterative(testCase.head, testCase.value, testCase.index)
	}
}

func BenchmarkIterative000(b *testing.B) { benchmarkIterative(b, testCases[0]) }
func BenchmarkIterative001(b *testing.B) { benchmarkIterative(b, testCases[1]) }
func BenchmarkIterative002(b *testing.B) { benchmarkIterative(b, testCases[2]) }
func BenchmarkIterative003(b *testing.B) { benchmarkIterative(b, testCases[3]) }
