package uncompress

import (
	"fmt"
	"strconv"
	"strings"
)

//nolint:varnamelen
func findAllStringSubmatchGenerator(s string) chan struct {
	number string
	char   int32
} {
	//nolint:varnamelen
	c := make(chan struct {
		number string
		char   int32
	})

	go func() {
		var number string

		for _, elem := range s {
			if (elem >= '0') && (elem <= '9') {
				number += string(elem)
			} else {
				char := elem
				c <- struct {
					number string
					char   int32
				}{number: number, char: char}
				number = ""
			}
		}

		close(c)
	}()

	return c
}

func Generator(s string) (string, error) {
	var result strings.Builder

	for match := range findAllStringSubmatchGenerator(s) {
		count, err := strconv.Atoi(match.number)
		if err != nil {
			return "", fmt.Errorf("failed to convert '%s' to int: %w", match.number, err)
		}

		result.WriteString(strings.Repeat(string(match.char), count))
	}

	return result.String(), nil
}
