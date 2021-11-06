package linkedlistcycle_test

import (
	"reflect"
	"testing"

	linkedlistcycle "github.com/ju-popov/structy.net/076-linked-list-cycle"
)

func TestWithSet(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := linkedlistcycle.WithSet(testCase.head)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkWithSet(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		linkedlistcycle.WithSet(testCase.head)
	}
}

func BenchmarkWithSet000(b *testing.B) { benchmarkWithSet(b, testCases[0]) }
func BenchmarkWithSet001(b *testing.B) { benchmarkWithSet(b, testCases[1]) }
func BenchmarkWithSet002(b *testing.B) { benchmarkWithSet(b, testCases[2]) }
func BenchmarkWithSet003(b *testing.B) { benchmarkWithSet(b, testCases[3]) }
func BenchmarkWithSet004(b *testing.B) { benchmarkWithSet(b, testCases[4]) }
func BenchmarkWithSet005(b *testing.B) { benchmarkWithSet(b, testCases[5]) }
