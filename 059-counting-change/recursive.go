package countingchange

type memoryKey struct {
	amount    int
	coinIndex int
}

func helper(amount int, coins []int, coinIndex int, memory map[memoryKey]int) int {
	if coinIndex >= len(coins) {
		return 0
	}

	key := memoryKey{amount: amount, coinIndex: coinIndex}
	if value, ok := memory[key]; ok {
		return value
	}

	count := 0
	coin := coins[coinIndex]

	for quantity := 0; quantity*coin <= amount; quantity++ {
		remainder := amount - quantity*coin
		if remainder == 0 {
			count++
		} else {
			count += helper(remainder, coins, coinIndex+1, memory)
		}
	}

	memory[key] = count

	return memory[key]
}

func Recursive(amount int, coins []int) int {
	memory := make(map[memoryKey]int)

	return helper(amount, coins, 0, memory)
}
