package uncompress

import (
	"regexp"
	"strconv"
	"strings"
)

func Regex(s string) string {
	var result strings.Builder

	re := regexp.MustCompile(`(\d+)(.)`)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		number, _ := strconv.Atoi(match[1])
		char := match[2]
		result.WriteString(strings.Repeat(char, number))
	}

	return result.String()
}
