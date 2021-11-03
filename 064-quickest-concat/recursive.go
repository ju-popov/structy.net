package canconcat

import "strings"

func helper(s string, words []string, memory map[string]int) int {
	if value, ok := memory[s]; ok {
		return value
	}

	if s == "" {
		return 0
	}

	minResult := -1
	for _, word := range words {
		if strings.HasPrefix(s, word) {
			if result := helper(s[len(word):], words, memory); result >= 0 {
				if (minResult == -1) || (result < minResult) {
					minResult = result
				}
			}
		}
	}

	if minResult >= 0 {
		memory[s] = minResult + 1
	} else {
		memory[s] = -1
	}

	return memory[s]
}

func Recursive(s string, words []string) int {
	memory := make(map[string]int)

	return helper(s, words, memory)
}
