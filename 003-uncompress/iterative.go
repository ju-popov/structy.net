package uncompress

import (
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

	number := 0

	for _, elem := range s {
		if (elem >= '0') && (elem <= '9') {
			//nolint:gomnd
			number = number*10 + int(elem-'0')
		} else {
			result = append(result, struct {
				number int
				char   string
			}{number: number, char: string(elem)})
			number = 0
		}
	}

	return result
}

func Iterative(s string) string {
	parts := make([]string, 0)

	for _, match := range findAllStringSubmatch(s) {
		parts = append(parts, strings.Repeat(match.char, match.number))
	}

	return strings.Join(parts, "")
}
