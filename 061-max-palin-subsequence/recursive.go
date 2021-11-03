package maxpalinsubsequence

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

	firstChar := str[0]
	lastChar := str[len(str)-1]

	if firstChar == lastChar {
		removeFirstAndLast := helper(str[1:len(str)-1], memory)
		memory[str] = removeFirstAndLast + 2
	} else {
		removeFirst := helper(str[1:], memory)
		removeLast := helper(str[:len(str)-1], memory)
		if removeFirst > removeLast {
			memory[str] = removeFirst
		} else {
			memory[str] = removeLast
		}
	}

	return memory[str]
}

func Recursive(str string) int {
	memory := make(map[string]int)

	return helper(str, memory)
}
