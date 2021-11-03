package summingsquares

func helper(n int, perfectSquares []int, memory map[int]int) int {
	if value, ok := memory[n]; ok {
		return value
	}

	if n == 0 {
		return 0
	}

	minResult := -1

	for _, perfectSquare := range perfectSquares {
		if n < perfectSquare {
			break
		}

		result := helper(n-perfectSquare, perfectSquares, memory)
		if (minResult == -1) || (result < minResult) {
			minResult = result
		}
	}

	// minResult is always >= 0
	memory[n] = minResult + 1

	return memory[n]
}

func calcPerfectSquares(max int) []int {
	result := make([]int, 0)
	for i := 1; i*i <= max; i++ {
		result = append(result, i*i)
	}

	return result
}

func Recursive(n int) int {
	perfectSquares := calcPerfectSquares(n)
	memory := make(map[int]int)

	return helper(n, perfectSquares, memory)
}
