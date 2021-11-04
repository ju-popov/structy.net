package sumpossible_test

import (
	"reflect"
	"testing"

	sumpossible "github.com/ju-popov/structy.net/053-sum-possible"
)

func TestDynamicProgramming(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := sumpossible.DynamicProgramming(testCase.amount, testCase.numbers)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkDynamicProgramming(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		sumpossible.DynamicProgramming(testCase.amount, testCase.numbers)
	}
}

func BenchmarkDynamicProgramming000(b *testing.B) { benchmarkDynamicProgramming(b, testCases[0]) }
func BenchmarkDynamicProgramming001(b *testing.B) { benchmarkDynamicProgramming(b, testCases[1]) }
func BenchmarkDynamicProgramming002(b *testing.B) { benchmarkDynamicProgramming(b, testCases[2]) }
func BenchmarkDynamicProgramming003(b *testing.B) { benchmarkDynamicProgramming(b, testCases[3]) }
func BenchmarkDynamicProgramming004(b *testing.B) { benchmarkDynamicProgramming(b, testCases[4]) }
func BenchmarkDynamicProgramming005(b *testing.B) { benchmarkDynamicProgramming(b, testCases[5]) }
func BenchmarkDynamicProgramming006(b *testing.B) { benchmarkDynamicProgramming(b, testCases[6]) }
func BenchmarkDynamicProgramming007(b *testing.B) { benchmarkDynamicProgramming(b, testCases[7]) }
func BenchmarkDynamicProgramming008(b *testing.B) { benchmarkDynamicProgramming(b, testCases[8]) }
