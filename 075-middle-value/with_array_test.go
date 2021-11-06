package middlevalue_test

import (
	"reflect"
	"testing"

	middlevalue "github.com/ju-popov/structy.net/075-middle-value"
)

func TestWithArray(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := middlevalue.WithArray(testCase.head)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkWithArray(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		middlevalue.WithArray(testCase.head)
	}
}

func BenchmarkWithArray000(b *testing.B) { benchmarkWithArray(b, testCases[0]) }
func BenchmarkWithArray001(b *testing.B) { benchmarkWithArray(b, testCases[1]) }
func BenchmarkWithArray002(b *testing.B) { benchmarkWithArray(b, testCases[2]) }
func BenchmarkWithArray003(b *testing.B) { benchmarkWithArray(b, testCases[3]) }
func BenchmarkWithArray004(b *testing.B) { benchmarkWithArray(b, testCases[4]) }
