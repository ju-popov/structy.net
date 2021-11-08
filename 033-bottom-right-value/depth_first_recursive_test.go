package bottomrightvalue_test

import (
	"reflect"
	"testing"

	bottomrightvalue "github.com/ju-popov/structy.net/033-bottom-right-value"
)

func TestDepthFirstRecusrsive(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			actual := bottomrightvalue.DepthFirstRecusrsive(testCase.root)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDepthFirstRecusrsive(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		bottomrightvalue.DepthFirstRecusrsive(testCase.root)
	}
}

func BenchmarkDepthFirstRecusrsive000(b *testing.B) { benchmarkDepthFirstRecusrsive(b, testCases[0]) }
func BenchmarkDepthFirstRecusrsive001(b *testing.B) { benchmarkDepthFirstRecusrsive(b, testCases[1]) }
func BenchmarkDepthFirstRecusrsive002(b *testing.B) { benchmarkDepthFirstRecusrsive(b, testCases[2]) }
func BenchmarkDepthFirstRecusrsive003(b *testing.B) { benchmarkDepthFirstRecusrsive(b, testCases[3]) }
func BenchmarkDepthFirstRecusrsive004(b *testing.B) { benchmarkDepthFirstRecusrsive(b, testCases[4]) }
