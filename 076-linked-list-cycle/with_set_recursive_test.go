package linkedlistcycle_test

import (
	"reflect"
	"testing"

	linkedlistcycle "github.com/ju-popov/structy.net/076-linked-list-cycle"
)

func TestWithSetRecursive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := linkedlistcycle.WithSetRecursive(testCase.head)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkWithSetRecursive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		linkedlistcycle.WithSetRecursive(testCase.head)
	}
}

func BenchmarkWithSetRecursive000(b *testing.B) { benchmarkWithSetRecursive(b, testCases[0]) }
func BenchmarkWithSetRecursive001(b *testing.B) { benchmarkWithSetRecursive(b, testCases[1]) }
func BenchmarkWithSetRecursive002(b *testing.B) { benchmarkWithSetRecursive(b, testCases[2]) }
func BenchmarkWithSetRecursive003(b *testing.B) { benchmarkWithSetRecursive(b, testCases[3]) }
func BenchmarkWithSetRecursive004(b *testing.B) { benchmarkWithSetRecursive(b, testCases[4]) }
func BenchmarkWithSetRecursive005(b *testing.B) { benchmarkWithSetRecursive(b, testCases[5]) }
