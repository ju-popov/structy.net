package maxpathsum

func helper(grid [][]int, y int, x int, memory [][]int) int {
	if y >= len(grid) {
		return 0
	}

	if x >= len(grid[y]) {
		return 0
	}

	if value := memory[y][x]; value > -1 {
		return value
	}

	memory[y][x] = grid[y][x]

	down := helper(grid, y+1, x, memory)
	right := helper(grid, y, x+1, memory)

	if down > right {
		memory[y][x] += down
	} else {
		memory[y][x] += right
	}

	return memory[y][x]
}

func Recursive(grid [][]int) int {
	memory := make([][]int, len((grid)))
	for y, row := range grid {
		memory[y] = make([]int, len(row))
		for x := range row {
			memory[y][x] = -1
		}
	}

	return helper(grid, 0, 0, memory)
}
