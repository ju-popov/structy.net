package bestbridge

const water = "W"

type queueValue struct {
	y        int
	x        int
	distance int
}

type visitedKey struct {
	y int
	x int
}

func neighborOffsets() [4][2]int {
	return [4][2]int{
		{-1, 0}, // top
		{+1, 0}, // down
		{0, -1}, // left
		{0, +1}, // right
	}
}

func inRange(grid [][]string, y int, x int) bool {
	if (y < 0) || (y >= len(grid)) {
		return false
	}

	if (x < 0) || (x >= len(grid[y])) {
		return false
	}

	return true
}

//nolint:varnamelen
func traverseIsland(grid [][]string, y int, x int, visited map[visitedKey]bool) bool {
	if grid[y][x] == water {
		return false
	}

	key := visitedKey{y: y, x: x}

	visited[key] = true

	for _, offset := range neighborOffsets() {
		neighborY := y + offset[0]
		neighborX := x + offset[1]
		neighborKey := visitedKey{y: neighborY, x: neighborX}

		if inRange(grid, neighborY, neighborX) && !visited[neighborKey] {
			traverseIsland(grid, neighborY, neighborX, visited)
		}
	}

	return true
}

func BreadthFirst(grid [][]string) int {
	visited := make(map[visitedKey]bool)

	// Look for and mark an island
out:
	for y, row := range grid {
		for x := range row {
			if traverseIsland(grid, y, x, visited) {
				break out
			}
		}
	}

	// Initial queue
	queue := make([]queueValue, 0)

	for key := range visited {
		queue = append(queue, queueValue{y: key.y, x: key.x, distance: 0})
	}

	// Breadth first
	for len(queue) > 0 {
		//nolint:varnamelen
		y := queue[0].y
		//nolint:varnamelen
		x := queue[0].x
		distance := queue[0].distance
		queue = queue[1:]

		key := visitedKey{y: y, x: x}

		if !visited[key] && (grid[y][x] != water) {
			return distance - 1
		}

		visited[key] = true

		for _, neighborOffset := range neighborOffsets() {
			neighborY := y + neighborOffset[0]
			neighborX := x + neighborOffset[1]
			neighborDistance := distance + 1
			neighborKey := visitedKey{y: neighborY, x: neighborX}

			if inRange(grid, neighborY, neighborX) && !visited[neighborKey] {
				// Helpful: Premature return
				if grid[neighborY][neighborX] != water {
					return neighborDistance - 1
				}

				queue = append(queue, queueValue{y: neighborY, x: neighborX, distance: neighborDistance})
			}
		}
	}

	return -1
}
