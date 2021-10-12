package mostfrequentchar

func OnePass(input string) string {
	chars := map[int32]struct {
		count int64
		index int
	}{}

	var (
		mostFrequentChar int32
	)

	for index, element := range input {
		if val, ok := chars[element]; ok {
			chars[element] = struct {
				count int64
				index int
			}{count: val.count + 1, index: val.index}
		} else {
			chars[element] = struct {
				count int64
				index int
			}{count: 1, index: index}
		}

		if chars[element].count == chars[mostFrequentChar].count {
			if chars[element].index < chars[mostFrequentChar].index {
				mostFrequentChar = element
			}
		} else if chars[element].count > chars[mostFrequentChar].count {
			mostFrequentChar = element
		}
	}

	return string(mostFrequentChar)
}
