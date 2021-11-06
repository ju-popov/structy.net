package substitutesynonyms

import "strings"

func helper(words []string, synonyms map[string][]string) []string {
	results := make([]string, 0)

	if len(words) == 0 {
		results = append(results, "")
		return results
	}

	firstWord := words[0]
	otherWords := words[1:]

	replaceWords := []string{firstWord}
	if _, ok := synonyms[firstWord]; ok {
		replaceWords = synonyms[firstWord]
	}

	for _, replaceWord := range replaceWords {
		for _, substituted := range helper(otherWords, synonyms) {
			if substituted == "" {
				results = append(results, replaceWord)
			} else {
				results = append(results, replaceWord+" "+substituted)
			}
		}
	}

	return results
}

func Recursive(sentance string, synonyms map[string][]string) []string {
	return helper(strings.Split(sentance, " "), synonyms)
}
