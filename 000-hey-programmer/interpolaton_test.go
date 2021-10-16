package heyprogrammer_test

import (
	"reflect"
	"testing"

	heyprogrammer "github.com/ju-popov/structy.net/000-hey-programmer"
)

func TestInterpolation(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := heyprogrammer.Interpolation(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input: '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkInterpolation(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		heyprogrammer.Interpolation(tc.input)
	}
}

func BenchmarkInterpolation000(b *testing.B) { benchmarkInterpolation(b, testCases[0]) }
func BenchmarkInterpolation001(b *testing.B) { benchmarkInterpolation(b, testCases[1]) }
func BenchmarkInterpolation002(b *testing.B) { benchmarkInterpolation(b, testCases[2]) }
