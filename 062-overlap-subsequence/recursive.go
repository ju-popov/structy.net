package overlapsubsequence

type memoryKey struct {
	str1 string
	str2 string
}

func maxInt(values ...int) int {
	maxValue := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	return maxValue
}

func helper(str1 string, str2 string, memory map[memoryKey]int) int {
	key := memoryKey{str1: str1, str2: str2}
	if value, ok := memory[key]; ok {
		return value
	}

	if (str1 == "") || (str2 == "") {
		return 0
	}

	if str1[0] == str2[0] {
		memory[key] = helper(str1[1:], str2[1:], memory) + 1
	} else {
		removeFirstCharStr1 := helper(str1[1:], str2, memory)
		removeFirstCharStr2 := helper(str1, str2[1:], memory)
		memory[key] = maxInt(removeFirstCharStr1, removeFirstCharStr2)
	}

	return memory[key]
}

func Recursive(str1 string, str2 string) int {
	memory := make(map[memoryKey]int)

	return helper(str1, str2, memory)
}
