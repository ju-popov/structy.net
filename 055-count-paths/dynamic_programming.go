package countpaths

const wall = "X"

type memoryKey struct {
	y int
	x int
}

//nolint:varnamelen
func helper(grid [][]string, y int, x int, memory map[memoryKey]int) int {
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

	if grid[y][x] == wall {
		return 0
	}

	if (y == len(grid)-1) && (x == len(grid[y])-1) {
		memory[key] = 1
	} else {
		memory[key] = helper(grid, y+1, x, memory) + helper(grid, y, x+1, memory)
	}

	return memory[key]
}

func DynamicProgramming(grid [][]string) int {
	memory := make(map[memoryKey]int)

	return helper(grid, 0, 0, memory)
}
