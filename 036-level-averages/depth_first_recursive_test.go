package levelaverages_test

import (
	"reflect"
	"testing"

	levelaverages "github.com/ju-popov/structy.net/036-level-averages"
)

func TestDepthFirstRecursive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := levelaverages.DepthFirstRecursive(testCase.root)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstRecursive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		levelaverages.DepthFirstRecursive(testCase.root)
	}
}

func BenchmarkDepthFirstRecursive000(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[0]) }
func BenchmarkDepthFirstRecursive001(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[1]) }
func BenchmarkDepthFirstRecursive002(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[2]) }
func BenchmarkDepthFirstRecursive003(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[3]) }
func BenchmarkDepthFirstRecursive004(b *testing.B) { benchmarkDepthFirstRecursive(b, testCases[4]) }
