package mostfrequentchar

func OnePass(input string) string {
	charsCount := map[int32]int64{}
	charsIndex := map[int32]int{}

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
