package substitutesynonyms

import "strings"

func firstWordInSynonyms(words []string, synonyms map[string][]string) int {
	for index, word := range words {
		if _, ok := synonyms[word]; ok {
			return index
		}
	}

	return -1
}

func helper(sentances []string, synonyms map[string][]string) []string {
	results := make([]string, 0)

	for _, sentance := range sentances {
		words := strings.Split(sentance, " ")
		wordIndex := firstWordInSynonyms(words, synonyms)

		if wordIndex < 0 {
			results = append(results, sentance)
			continue
		}

		replaced := make([]string, 0)
		for _, value := range synonyms[words[wordIndex]] {
			words[wordIndex] = value
			replaced = append(replaced, strings.Join(words, " "))
		}

		results = append(results, helper(replaced, synonyms)...)
	}

	return results
}

func Recursive(sentance string, synonyms map[string][]string) []string {
	return helper([]string{sentance}, synonyms)
}
