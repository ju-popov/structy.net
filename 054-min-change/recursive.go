package minchange

func helper(amount int, coins []int, memory map[int]int) int {
	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	if value, ok := memory[amount]; ok {
		return value
	}

	minCoins := -1

	for _, number := range coins {
		if coinsCount := helper(amount-number, coins, memory); coinsCount >= 0 {
			if (minCoins == -1) || (coinsCount < minCoins) {
				minCoins = coinsCount
			}
		}
	}

	if minCoins >= 0 {
		memory[amount] = minCoins + 1
	} else {
		memory[amount] = -1
	}

	return memory[amount]
}

func Recursive(amount int, coins []int) int {
	memory := make(map[int]int)

	return helper(amount, coins, memory)
}
