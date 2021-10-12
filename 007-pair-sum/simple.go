package pairsum

func Simple(numbers []int64, targetSum int64) [2]int {
	numbersIndexes := map[int64]int{}

	for index, element := range numbers {
		if val, ok := numbersIndexes[targetSum-element]; ok {
			return [2]int{val, index}
		}

		if _, ok := numbersIndexes[element]; !ok {
			numbersIndexes[element] = index
		}
	}

	return [2]int{-1, -1}
}
