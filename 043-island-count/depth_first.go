package islandcount

type visitedKey struct {
	y int
	x int
}

//nolint:varnamelen
func explore(grid [][]string, y int, x int, visited map[visitedKey]bool) bool {
	if (y < 0) || (y >= len(grid)) {
		return false
	}

	row := grid[y]

	if (x < 0) || (x >= len(row)) {
		return false
	}

	if row[x] != "L" {
		return false
	}

	key := visitedKey{y: y, x: x}

	if visited[key] {
		return false
	}

	visited[key] = true

	explore(grid, y-1, x, visited)
	explore(grid, y+1, x, visited)
	explore(grid, y, x-1, visited)
	explore(grid, y, x+1, visited)

	return true
}

func DepthFirst(grid [][]string) int {
	result := 0

	visited := make(map[visitedKey]bool)

	for y, row := range grid {
		for x := range row {
			if explore(grid, y, x, visited) {
				result++
			}
		}
	}

	return result
}
