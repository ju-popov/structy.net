package islandcount

func explore(grid [][]string, y int, x int, visited [][]bool) bool {
	if (y < 0) || (y >= len(grid)) {
		return false
	}

	row := grid[y]

	if (x < 0) || (x >= len(row)) {
		return false
	}

	if (row[x] != "L") || visited[y][x] {
		return false
	}

	visited[y][x] = true

	explore(grid, y-1, x, visited)
	explore(grid, y+1, x, visited)
	explore(grid, y, x-1, visited)
	explore(grid, y, x+1, visited)

	return true
}

func DepthFirst(grid [][]string) int {
	result := 0

	visited := make([][]bool, len(grid))
	for index, row := range grid {
		visited[index] = make([]bool, len(row))
	}

	for y, row := range grid {
		for x := range row {
			if explore(grid, y, x, visited) {
				result++
			}
		}
	}

	return result
}
