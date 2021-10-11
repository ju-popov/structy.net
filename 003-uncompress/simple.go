package uncompress

import (
	"fmt"
	"strconv"
	"strings"
)

func findAllStringSubmatch(s string) [][3]string {
	var result [][3]string

	count := ""

	for _, elem := range s {
		if (elem >= '0') && (elem <= '9') {
			count += string(elem)
		} else {
			char := string(elem)
			result = append(result, [3]string{count + char, count, char})
			count = ""
		}
	}

	return result
}

func Simple(s string) (string, error) {
	var result strings.Builder

	for _, match := range findAllStringSubmatch(s) {
		count, err := strconv.Atoi(match[1])
		if err != nil {
			return "", fmt.Errorf("failed to convert '%s' to int: %w", match[1], err)
		}

		char := match[2]

		result.WriteString(strings.Repeat(char, count))
	}

	return result.String(), nil
}
