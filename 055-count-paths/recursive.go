package countpaths

const wall = "X"

func helper(grid [][]string, y int, x int, memory [][]int) int {
	if y >= len(grid) {
		return 0
	}

	if x >= len(grid[y]) {
		return 0
	}

	if value := memory[y][x]; value > -1 {
		return value
	}

	if grid[y][x] == wall {
		return 0
	}

	if (y == len(grid)-1) && (x == len(grid[y])-1) {
		memory[y][x] = 1
	} else {
		memory[y][x] = helper(grid, y+1, x, memory) + helper(grid, y, x+1, memory)
	}

	return memory[y][x]
}

func Recursive(grid [][]string) int {
	memory := make([][]int, len((grid)))
	for y, row := range grid {
		memory[y] = make([]int, len(row))
		for x := range row {
			memory[y][x] = -1
		}
	}

	return helper(grid, 0, 0, memory)
}
