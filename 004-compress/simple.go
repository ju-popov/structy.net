package compress

import (
	"strconv"
	"strings"
)

func findAllStringSubmatch(s string) []struct {
	number int64
	char   int32
} {
	var result []struct {
		number int64
		char   int32
	}

	var (
		count int64
		char  int32
	)

	for index, elem := range s {
		if index == 0 {
			char = elem
			count = 1
		} else {
			if elem == char {
				count++
			} else {
				result = append(result, struct {
					number int64
					char   int32
				}{number: count, char: char})
				char = elem
				count = 1
			}
		}
	}

	result = append(result, struct {
		number int64
		char   int32
	}{number: count, char: char})

	return result
}

func Simple(s string) (string, error) {
	var result strings.Builder

	for _, match := range findAllStringSubmatch(s) {
		if match.number == 1 {
			result.WriteRune(match.char)
		} else if match.number > 1 {
			//nolint:gomnd
			result.WriteString(strconv.FormatInt(match.number, 10))
			result.WriteRune(match.char)
		}
	}

	return result.String(), nil
}
