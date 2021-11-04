package canconcat

import "strings"

//nolint:varnamelen
func helper(s string, words []string, memory map[string]bool) bool {
	if value, ok := memory[s]; ok {
		return value
	}

	if s == "" {
		return true
	}

	for _, word := range words {
		if strings.HasPrefix(s, word) && helper(s[len(word):], words, memory) {
			memory[s] = true

			return memory[s]
		}
	}

	memory[s] = false

	return memory[s]
}

func DynamicProgramming(s string, words []string) bool {
	memory := make(map[string]bool)

	return helper(s, words, memory)
}
