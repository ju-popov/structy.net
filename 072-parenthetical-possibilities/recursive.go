package parentheticalpossibilities

import (
	"strings"
)

//nolint:varnamelen
func helper(ss []string) []string {
	results := make([]string, 0)

	for _, s := range ss {
		startIndex := strings.Index(s, "(")
		endIndex := strings.Index(s, ")")

		if (startIndex < 0) || (endIndex < 0) {
			results = append(results, s)

			continue
		}

		uncompressed := make([]string, 0)
		for index := startIndex + 1; index < endIndex; index++ {
			uncompressed = append(uncompressed, s[:startIndex]+string(s[index])+s[endIndex+1:])
		}

		results = append(results, helper(uncompressed)...)
	}

	return results
}

func Recursive(s string) []string {
	return helper([]string{s})
}
