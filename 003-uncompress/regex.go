package uncompress

import (
	"regexp"
	"strconv"
	"strings"
)

func Regex(s string) string {
	parts := make([]string, 0)

	re := regexp.MustCompile(`(\d+)(.)`)
	for _, match := range re.FindAllStringSubmatch(s, -1) {
		number, _ := strconv.Atoi(match[1])
		parts = append(parts, strings.Repeat(match[2], number))
	}

	return strings.Join(parts, "")
}
