package compress

import (
	"fmt"
	"strings"
)

func findAllStringSubmatch(s string) []struct {
	count int64
	char  string
} {
	var result []struct {
		count int64
		char  string
	}

	count := int64(0)
	char := ""

	for _, elem := range s {
		if string(elem) == char {
			count++
		} else {
			result = append(result, struct {
				count int64
				char  string
			}{count: count, char: char})
			char = string(elem)
			count = 1
		}
	}

	result = append(result, struct {
		count int64
		char  string
	}{count: count, char: char})

	return result
}

func Simple(s string) (string, error) {
	var result strings.Builder

	for _, match := range findAllStringSubmatch(s) {
		if match.count == 1 {
			result.WriteString(match.char)
		} else if match.count > 1 {
			result.WriteString(fmt.Sprintf("%d%s", match.count, match.char))
		}
	}

	return result.String(), nil
}
