package fib_test

import (
	"reflect"
	"testing"

	fib "github.com/ju-popov/structy.net/051-fib"
)

func TestRecursiveWithMemory(t *testing.T) {
	t.Parallel()

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			actual := fib.RecursiveWithMemory(tc.n)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for name: '%v' is: '%v', but the actual result is: '%v'", tc.name, tc.expected, actual)
			}
		})
	}
}

func benchmarkRecursiveWithMemory(b *testing.B, tc testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		fib.RecursiveWithMemory(tc.n)
	}
}

func BenchmarkRecursiveWithMemory000(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[0]) }
func BenchmarkRecursiveWithMemory001(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[1]) }
func BenchmarkRecursiveWithMemory002(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[2]) }
func BenchmarkRecursiveWithMemory003(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[3]) }
func BenchmarkRecursiveWithMemory004(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[4]) }
func BenchmarkRecursiveWithMemory005(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[5]) }
func BenchmarkRecursiveWithMemory006(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[6]) }
func BenchmarkRecursiveWithMemory007(b *testing.B) { benchmarkRecursiveWithMemory(b, testCases[7]) }
