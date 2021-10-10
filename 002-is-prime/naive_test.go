package isprime_test

import (
	"reflect"
	"testing"

	isprime "github.com/ju-popov/structy.net/002-is-prime"
)

type testCase struct {
	name     string
	input    int64
	expected bool
}

//nolint:gochecknoglobals
var testCases = []testCase{
	{
		name:     "test_00",
		input:    2,
		expected: true,
	},
	{
		name:     "test_01",
		input:    3,
		expected: true,
	},
	{
		name:     "test_02",
		input:    4,
		expected: false,
	},
	{
		name:     "test_03",
		input:    5,
		expected: true,
	},
	{
		name:     "test_04",
		input:    6,
		expected: false,
	},
	{
		name:     "test_05",
		input:    7,
		expected: true,
	},
	{
		name:     "test_06",
		input:    8,
		expected: false,
	},
	{
		name:     "test_07",
		input:    25,
		expected: false,
	},
	{
		name:     "test_08",
		input:    31,
		expected: true,
	},
	{
		name:     "test_09",
		input:    2017,
		expected: true,
	},
	{
		name:     "test_10",
		input:    2048,
		expected: false,
	},
	{
		name:     "test_11",
		input:    1,
		expected: false,
	},
}

func TestNaive(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := isprime.Naive(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}

func benchmarkNaive(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		isprime.Naive(tc.input)
	}
}

func BenchmarkNaive000(b *testing.B) { benchmarkNaive(b, testCases[0]) }
func BenchmarkNaive001(b *testing.B) { benchmarkNaive(b, testCases[1]) }
func BenchmarkNaive002(b *testing.B) { benchmarkNaive(b, testCases[2]) }
func BenchmarkNaive003(b *testing.B) { benchmarkNaive(b, testCases[3]) }
func BenchmarkNaive004(b *testing.B) { benchmarkNaive(b, testCases[4]) }
func BenchmarkNaive005(b *testing.B) { benchmarkNaive(b, testCases[5]) }
func BenchmarkNaive006(b *testing.B) { benchmarkNaive(b, testCases[6]) }
func BenchmarkNaive007(b *testing.B) { benchmarkNaive(b, testCases[7]) }
func BenchmarkNaive008(b *testing.B) { benchmarkNaive(b, testCases[8]) }
func BenchmarkNaive009(b *testing.B) { benchmarkNaive(b, testCases[9]) }
func BenchmarkNaive010(b *testing.B) { benchmarkNaive(b, testCases[10]) }
func BenchmarkNaive011(b *testing.B) { benchmarkNaive(b, testCases[11]) }
