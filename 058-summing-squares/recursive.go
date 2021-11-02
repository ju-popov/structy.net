package summingsquares

func helper(n int, squares []int, memory map[int]int) int {
	if n == 0 {
		return 0
	}

	if value, ok := memory[n]; ok {
		return value
	}

	minResult := -1

	for _, square := range squares {
		if n < square {
			break
		}

		if result := helper(n-square, squares, memory); (minResult == -1) || (result < minResult) {
			minResult = result
		}
	}

	memory[n] = minResult + 1

	return memory[n]
}

func findPerfectSquares(max int) []int {
	result := make([]int, 0)
	for i := 1; i*i <= max; i++ {
		result = append(result, i*i)
	}

	return result
}

func Recursive(n int) int {
	squares := findPerfectSquares(n)
	memory := make(map[int]int)

	return helper(n, squares, memory)
}
