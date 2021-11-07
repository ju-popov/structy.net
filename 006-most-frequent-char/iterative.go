package mostfrequentchar

//nolint:varnamelen
func Iterative(s string) string {
	chars := make(map[int32]int)

	for _, element := range s {
		chars[element]++
	}

	var mostFrequentChar int32
	for _, element := range s {
		if (mostFrequentChar == 0) || (chars[element] > chars[mostFrequentChar]) {
			mostFrequentChar = element
		}
	}

	return string(mostFrequentChar)
}
