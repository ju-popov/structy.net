package pairproduct

func Iterative(numbers []int64, targetProduct int64) [2]int {
	numbersIndexes := map[int64]int{}

	for index, element := range numbers {
		if targetProduct%element == 0 {
			if val, ok := numbersIndexes[targetProduct/element]; ok {
				return [2]int{val, index}
			}
		}

		if _, ok := numbersIndexes[element]; !ok {
			numbersIndexes[element] = index
		}
	}

	return [2]int{-1, -1}
}
