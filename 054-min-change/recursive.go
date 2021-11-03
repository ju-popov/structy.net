package minchange

func helper(amount int, coins []int, memory map[int]int) int {
	if value, ok := memory[amount]; ok {
		return value
	}

	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	minResult := -1

	for _, number := range coins {
		if result := helper(amount-number, coins, memory); result >= 0 {
			if (minResult == -1) || (result < minResult) {
				minResult = result
			}
		}
	}

	if minResult >= 0 {
		memory[amount] = minResult + 1
	} else {
		memory[amount] = -1
	}

	return memory[amount]
}

func Recursive(amount int, coins []int) int {
	memory := make(map[int]int)

	return helper(amount, coins, memory)
}
