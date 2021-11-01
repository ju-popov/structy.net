package bestbridge

const water = "W"

type cellData struct {
	y        int
	x        int
	distance int
}

func offsets() [4][2]int {
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

func traverseIsland(grid [][]string, y int, x int, visited [][]bool) bool {
	if grid[y][x] == water {
		return false
	}

	visited[y][x] = true

	for _, offset := range offsets() {
		nextY, nextX := y+offset[0], x+offset[1]
		if inRange(grid, nextY, nextX) && !visited[nextY][nextX] {
			traverseIsland(grid, nextY, nextX, visited)
		}
	}

	return true
}

func BreadthFirst(grid [][]string) int {
	// Visited
	visited := make([][]bool, len(grid))
	for y, row := range grid {
		visited[y] = make([]bool, len(row))
	}

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
	queue := make([]cellData, 0)

	for y, row := range visited {
		for x, value := range row {
			if value {
				queue = append(queue, cellData{y: y, x: x, distance: 0})
			}
		}
	}

	// Breadth first
	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		if !visited[cell.y][cell.x] && (grid[cell.y][cell.x] != water) {
			return cell.distance - 1
		}

		visited[cell.y][cell.x] = true

		for _, offset := range offsets() {
			nextY, nextX, nextDistance := cell.y+offset[0], cell.x+offset[1], cell.distance+1

			if inRange(grid, nextY, nextX) && !visited[nextY][nextX] {
				// Helpful: Premature return
				if grid[nextY][nextX] != water {
					return nextDistance - 1
				}

				queue = append(queue, cellData{y: nextY, x: nextX, distance: nextDistance})
			}
		}
	}

	return -1
}
