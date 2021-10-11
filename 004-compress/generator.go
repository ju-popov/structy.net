package compress

import (
	"fmt"
	"strings"
)

func findAllStringSubmatchGenerator(s string) chan struct {
	count int64
	char  string
} {
	c := make(chan struct {
		count int64
		char  string
	})

	go func() {
		count := int64(0)
		char := ""

		for _, elem := range s {
			if string(elem) == char {
				count++
			} else {
				c <- struct {
					count int64
					char  string
				}{count: count, char: char}
				char = string(elem)
				count = 1
			}
		}

		c <- struct {
			count int64
			char  string
		}{count: count, char: char}

		close(c)
	}()

	return c
}

func Generator(s string) (string, error) {
	var result strings.Builder

	for match := range findAllStringSubmatchGenerator(s) {
		if match.count == 1 {
			result.WriteString(match.char)
		} else if match.count > 1 {
			result.WriteString(fmt.Sprintf("%d%s", match.count, match.char))
		}
	}

	return result.String(), nil
}
