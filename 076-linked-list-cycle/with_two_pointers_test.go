package linkedlistcycle_test

import (
	"reflect"
	"testing"

	linkedlistcycle "github.com/ju-popov/structy.net/076-linked-list-cycle"
)

func TestWithTwoPointers(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := linkedlistcycle.WithTwoPointers(testCase.head)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkWithTwoPointers(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		linkedlistcycle.WithTwoPointers(testCase.head)
	}
}

func BenchmarkWithTwoPointers000(b *testing.B) { benchmarkWithTwoPointers(b, testCases[0]) }
func BenchmarkWithTwoPointers001(b *testing.B) { benchmarkWithTwoPointers(b, testCases[1]) }
func BenchmarkWithTwoPointers002(b *testing.B) { benchmarkWithTwoPointers(b, testCases[2]) }
func BenchmarkWithTwoPointers003(b *testing.B) { benchmarkWithTwoPointers(b, testCases[3]) }
func BenchmarkWithTwoPointers004(b *testing.B) { benchmarkWithTwoPointers(b, testCases[4]) }
func BenchmarkWithTwoPointers005(b *testing.B) { benchmarkWithTwoPointers(b, testCases[5]) }
