package minimumisland

type visitedKey struct {
	y int
	x int
}

//nolint:varnamelen
func explore(grid [][]string, y int, x int, visited map[visitedKey]bool) int {
	if (y < 0) || (y >= len(grid)) {
		return 0
	}

	row := grid[y]

	if (x < 0) || (x >= len(row)) {
		return 0
	}

	if row[x] != "L" {
		return 0
	}

	key := visitedKey{y: y, x: x}

	if visited[key] {
		return 0
	}

	visited[key] = true

	result := 1

	result += explore(grid, y-1, x, visited)
	result += explore(grid, y+1, x, visited)
	result += explore(grid, y, x-1, visited)
	result += explore(grid, y, x+1, visited)

	return result
}

func DepthFirst(grid [][]string) int {
	minResult := 0

	visited := make(map[visitedKey]bool)

	for y, row := range grid {
		for x := range row {
			if result := explore(grid, y, x, visited); result > 0 {
				if (minResult == 0) || (result < minResult) {
					minResult = result
				}
			}
		}
	}

	return minResult
}
