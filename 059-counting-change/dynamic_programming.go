package countingchange

type memoryKey struct {
	amount    int
	coinIndex int
}

func helper(amount int, coins []int, coinIndex int, memory map[memoryKey]int) int {
	key := memoryKey{amount: amount, coinIndex: coinIndex}
	if value, ok := memory[key]; ok {
		return value
	}

	if coinIndex >= len(coins) {
		return 0
	}

	count := 0
	coin := coins[coinIndex]

	for qty := 0; (qty * coin) <= amount; qty++ {
		remainder := amount - qty*coin
		if remainder == 0 {
			count++
		} else {
			count += helper(remainder, coins, coinIndex+1, memory)
		}
	}

	memory[key] = count

	return memory[key]
}

func DynamicProgramming(amount int, coins []int) int {
	memory := make(map[memoryKey]int)

	return helper(amount, coins, 0, memory)
}
