package uncompress

import (
	"fmt"
	"strconv"
	"strings"
)

func findAllStringSubmatchGenerator(s string) chan [3]string {
	c := make(chan [3]string)

	go func() {
		count := ""

		for _, elem := range s {
			if (elem >= '0') && (elem <= '9') {
				count += string(elem)
			} else {
				char := string(elem)
				c <- [3]string{count + char, count, char}
				count = ""
			}
		}

		close(c)
	}()

	return c
}

func Generator(s string) (string, error) {
	var result strings.Builder

	for match := range findAllStringSubmatchGenerator(s) {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			return "", fmt.Errorf("failed to convert '%s' to int: %w", match[1], err)
		}

		char := match[2]

		result.WriteString(strings.Repeat(char, count))
	}

	return result.String(), nil
}
