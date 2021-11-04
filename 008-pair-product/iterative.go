package pairproduct

func Iterative(numbers []int, targetProduct int) [2]int {
	numbersIndexes := make(map[int]int)

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
