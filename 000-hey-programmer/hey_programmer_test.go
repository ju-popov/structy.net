package _000_hey_programmer

import (
	"reflect"
	"testing"
)

var testCases = []struct {
	name     string
	input    string
	expected string
}{
	{
		name: "alvin",
		input: "alvin",
		expected: "hey alvin",
	},
	{
		name: "jason",
		input: "jason",
		expected: "hey jason",
	},
	{
		name: "how now brown cow",
		input: "how now brown cow",
		expected: "hey how now brown cow",
	},
}

func TestHeyProgrammer(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := HeyProgrammer(tc.input)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Expected result for input '%s' is: '%v', but the actual result is: '%v'", tc.input, tc.expected, actual)
			}
		})
	}
}
