package sumpossible

func helper(amount int, numbers []int, memory map[int]bool) bool {
	if value, ok := memory[amount]; ok {
		return value
	}

	if amount < 0 {
		return false
	}

	if amount == 0 {
		return true
	}

	for _, number := range numbers {
		if helper(amount-number, numbers, memory) {
			memory[amount] = true

			return memory[amount]
		}
	}

	memory[amount] = false

	return memory[amount]
}

func DynamicProgramming(amount int, numbers []int) bool {
	memory := make(map[int]bool)

	return helper(amount, numbers, memory)
}
