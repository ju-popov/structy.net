package _001_max_value

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	input    []float64
	expected float64
}{
	{
		name: "test_00",
		input: []float64{4, 7, 2, 8, 10, 9},
		expected: 10,
	},
	{
		name: "test_01",
		input: []float64{10, 5, 40, 40.3},
		expected: 40.3,
	},
	{
		name: "test_02",
		input: []float64{-5, -2, -1, -11},
		expected: -1,
	},
	{
		name: "test_03",
		input: []float64{42},
		expected: 42,
	},
	{
		name: "test_04",
		input: []float64{1000, 8},
		expected: 1000,
	},
	{
		name: "test_05",
		input: []float64{1000, 8, 9000},
		expected: 9000,
	},
	{
		name: "test_06",
		input: []float64{2, 5, 1, 1, 4},
		expected: 5,
	},
}

func TestMaxValue(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := MaxValue(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input '%v' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}
