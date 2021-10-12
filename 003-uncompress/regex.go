package uncompress

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Regex(s string) (string, error) {
	var result strings.Builder

	re := regexp.MustCompile(`(\d+)(.)`)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		number, err := strconv.Atoi(match[1])
		if err != nil {
			return "", fmt.Errorf("failed to convert '%s' to int: %w", match[1], err)
		}

		char := match[2]

		result.WriteString(strings.Repeat(char, number))
	}

	return result.String(), nil
}
