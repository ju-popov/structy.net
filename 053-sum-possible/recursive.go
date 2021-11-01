package sumpossible

func helper(amount int, numbers []int, memory map[int]bool) bool {
	if amount < 0 {
		return false
	}

	if amount == 0 {
		return true
	}

	if value, ok := memory[amount]; ok {
		return value
	}

	for _, number := range numbers {
		if helper(amount-number, numbers, memory) {
			return true
		}
	}

	memory[amount] = false

	return false
}

func Recursive(amount int, numbers []int) bool {
	memory := make(map[int]bool)

	return helper(amount, numbers, memory)
}
