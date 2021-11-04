package anagrams

//nolint:varnamelen
func Iterative(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	charsCount := make(map[int32]int)

	for _, element := range s1 {
		charsCount[element]++
	}

	for _, element := range s2 {
		charsCount[element]--
	}

	for _, count := range charsCount {
		if count != 0 {
			return false
		}
	}

	return true
}
