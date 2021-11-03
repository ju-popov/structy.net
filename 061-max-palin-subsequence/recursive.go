package maxpalinsubsequence

func maxInt(values ...int) int {
	maxValue := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	return maxValue
}

func helper(str string, memory map[string]int) int {
	if value, ok := memory[str]; ok {
		return value
	}

	if len(str) == 0 {
		return 0
	}

	if len(str) == 1 {
		return 1
	}

	if firstChar, lastChar := str[0], str[len(str)-1]; firstChar == lastChar {
		removeFirstAndLastChar := helper(str[1:len(str)-1], memory)
		//nolint:gomnd
		memory[str] = removeFirstAndLastChar + 2
	} else {
		removeFirstChar := helper(str[1:], memory)
		removeLastChar := helper(str[:len(str)-1], memory)
		memory[str] = maxInt(removeFirstChar, removeLastChar)
	}

	return memory[str]
}

func Recursive(str string) int {
	memory := make(map[string]int)

	return helper(str, memory)
}
