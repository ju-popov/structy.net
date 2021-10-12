package mostfrequentchar

func TwoPass(input string) string {
	chars := map[int32]int64{}

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
