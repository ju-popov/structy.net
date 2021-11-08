package compress

import (
	"fmt"
	"strings"
)

//nolint:varnamelen
func findAllStringSubmatch(s string) []struct {
	number int
	char   string
} {
	var result []struct {
		number int
		char   string
	}

	var (
		count    int
		lastElem rune
	)

	for _, elem := range s {
		if elem == lastElem {
			count++
		} else {
			if count > 0 {
				result = append(result, struct {
					number int
					char   string
				}{number: count, char: string(lastElem)})
			}

			lastElem = elem
			count = 1
		}
	}

	if count > 0 {
		result = append(result, struct {
			number int
			char   string
		}{number: count, char: string(lastElem)})
	}

	return result
}

func Iterative(s string) string {
	parts := make([]string, 0)

	for _, match := range findAllStringSubmatch(s) {
		if match.number == 1 {
			parts = append(parts, match.char)
		} else if match.number > 1 {
			parts = append(parts, fmt.Sprintf("%d%s", match.number, match.char))
		}
	}

	return strings.Join(parts, "")
}
