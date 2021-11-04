package mostfrequentchar

func Iterative(input string) string {
	chars := make(map[int32]int)

	for _, element := range input {
		chars[element]++
	}

	var mostFrequentChar int32
	for _, element := range input {
		if (mostFrequentChar == 0) || (chars[element] > chars[mostFrequentChar]) {
			mostFrequentChar = element
		}
	}

	return string(mostFrequentChar)
}
