package pairsum

func Iterative(numbers []int, targetSum int) [2]int {
	numbersIndexes := make(map[int]int)

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
