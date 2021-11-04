package depthfirstvalues_test

import (
	"reflect"
	"testing"

	depthfirstvalues "github.com/ju-popov/structy.net/024-depth-first-values"
)

func TestDepthFirst(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := depthfirstvalues.DepthFirst(testCase.root)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirst(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		depthfirstvalues.DepthFirst(testCase.root)
	}
}

func BenchmarkDepthFirst000(b *testing.B) { benchmarkDepthFirst(b, testCases[0]) }
func BenchmarkDepthFirst001(b *testing.B) { benchmarkDepthFirst(b, testCases[1]) }
func BenchmarkDepthFirst002(b *testing.B) { benchmarkDepthFirst(b, testCases[2]) }
func BenchmarkDepthFirst003(b *testing.B) { benchmarkDepthFirst(b, testCases[3]) }
func BenchmarkDepthFirst004(b *testing.B) { benchmarkDepthFirst(b, testCases[4]) }
