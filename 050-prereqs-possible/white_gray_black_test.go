package prereqspossible_test

import (
	"reflect"
	"testing"

	prereqspossible "github.com/ju-popov/structy.net/050-prereqs-possible"
)

func TestWhiteGreyBlack(t *testing.T) {
	t.Parallel()

	for _, testCase := range testCases {
		testCase := testCase
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()
			actual := prereqspossible.WhiteGreyBlack(testCase.numCourses, testCase.prereqs)
			if !reflect.DeepEqual(actual, testCase.expected) {
				t.Errorf("Expected result for test name: '%v' is: '%v', but the actual result is: '%v'", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func benchmarkWhiteGreyBlack(b *testing.B, testCase testCase) {
	b.Helper()

	for n := 0; n < b.N; n++ {
		prereqspossible.WhiteGreyBlack(testCase.numCourses, testCase.prereqs)
	}
}

func BenchmarkWhiteGreyBlack000(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[0]) }
func BenchmarkWhiteGreyBlack001(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[1]) }
func BenchmarkWhiteGreyBlack002(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[2]) }
func BenchmarkWhiteGreyBlack003(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[3]) }
func BenchmarkWhiteGreyBlack004(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[4]) }
func BenchmarkWhiteGreyBlack005(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[5]) }
func BenchmarkWhiteGreyBlack006(b *testing.B) { benchmarkWhiteGreyBlack(b, testCases[6]) }
