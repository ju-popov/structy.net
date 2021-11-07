package heyprogrammer_test

import (
	"reflect"
	"testing"

	heyprogrammer "github.com/ju-popov/structy.net/000-hey-programmer"
)

func TestInterpolation(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := heyprogrammer.Interpolation(testCase.s)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkInterpolation(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		heyprogrammer.Interpolation(testCase.s)
	}
}

func BenchmarkInterpolation000(b *testing.B) { benchmarkInterpolation(b, testCases[0]) }
func BenchmarkInterpolation001(b *testing.B) { benchmarkInterpolation(b, testCases[1]) }
func BenchmarkInterpolation002(b *testing.B) { benchmarkInterpolation(b, testCases[2]) }
