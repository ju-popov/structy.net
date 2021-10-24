package uncompress

import (
	"fmt"
	"strconv"
	"strings"
)

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

func Iterative(s string) (string, error) {
	var result strings.Builder

	for _, match := range findAllStringSubmatch(s) {
		count, err := strconv.Atoi(match.number)
		if err != nil {
			return "", fmt.Errorf("failed to convert '%s' to int: %w", match.number, err)
		}

		result.WriteString(strings.Repeat(string(match.char), count))
	}

	return result.String(), nil
}
