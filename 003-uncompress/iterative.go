package uncompress

import (
	"strconv"
	"strings"
)

//nolint:varnamelen
func findAllStringSubmatch(s string) []struct {
	number string
	char   int32
} {
	var result []struct {
		number string
		char   int32
	}

	var number string

	for _, elem := range s {
		if (elem >= '0') && (elem <= '9') {
			number += string(elem)
		} else {
			char := elem
			result = append(result, struct {
				number string
				char   int32
			}{number: number, char: char})
			number = ""
		}
	}

	return result
}

func Iterative(s string) string {
	var result strings.Builder

	for _, match := range findAllStringSubmatch(s) {
		count, _ := strconv.Atoi(match.number)
		result.WriteString(strings.Repeat(string(match.char), count))
	}

	return result.String()
}
