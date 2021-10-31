package minimumisland

func explore(grid [][]string, y int, x int, visited [][]bool) int {
	if (y < 0) || (y >= len(grid)) {
		return 0
	}

	row := grid[y]

	if (x < 0) || (x >= len(row)) {
		return 0
	}

	if (row[x] != "L") || visited[y][x] {
		return 0
	}

	visited[y][x] = true

	result := 1

	result += explore(grid, y-1, x, visited)
	result += explore(grid, y+1, x, visited)
	result += explore(grid, y, x-1, visited)
	result += explore(grid, y, x+1, visited)

	return result
}

func DepthFirst(grid [][]string) int {
	minResult := 0

	visited := make([][]bool, len(grid))
	for index, row := range grid {
		visited[index] = make([]bool, len(row))
	}

	for y, row := range grid {
		for x := range row {
			if result := explore(grid, y, x, visited); (minResult == 0) || ((result > 0) && (result < minResult)) {
				minResult = result
			}
		}
	}

	return minResult
}
