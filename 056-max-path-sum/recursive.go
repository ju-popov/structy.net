package maxpathsum

type memoryKey struct {
	y int
	x int
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

func helper(grid [][]int, y int, x int, memory map[memoryKey]int) int {
	key := memoryKey{y: y, x: x}
	if value, ok := memory[key]; ok {
		return value
	}

	if y >= len(grid) {
		return 0
	}

	if x >= len(grid[y]) {
		return 0
	}

	down := helper(grid, y+1, x, memory)
	right := helper(grid, y, x+1, memory)

	memory[key] = grid[y][x] + maxInt(down, right)

	return memory[key]
}

func Recursive(grid [][]int) int {
	memory := make(map[memoryKey]int)

	return helper(grid, 0, 0, memory)
}
