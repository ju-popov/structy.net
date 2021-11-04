package mostfrequentchar

func OnePass(input string) string {
	charsCount := make(map[int32]int)
	charsIndex := make(map[int32]int)

	var mostFrequentChar int32

	for index, element := range input {
		charsCount[element]++

		if _, ok := charsIndex[element]; !ok {
			charsIndex[element] = index
		}

		if (mostFrequentChar == 0) ||
			(charsCount[element] > charsCount[mostFrequentChar]) ||
			((charsCount[element] == charsCount[mostFrequentChar]) && (charsIndex[element] < charsIndex[mostFrequentChar])) {
			mostFrequentChar = element
		}
	}

	return string(mostFrequentChar)
}
