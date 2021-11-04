package closestcarrot

type queueValue struct {
	y        int
	x        int
	distance int
}

type visitedKey struct {
	y int
	x int
}

const (
	wall   = "X"
	carrot = "C"
)

func offsets() [4][2]int {
	return [4][2]int{
		{-1, 0}, // top
		{+1, 0}, // down
		{0, -1}, // left
		{0, +1}, // right
	}
}

func BreadthFirst(grid [][]string, startY int, startX int) int {
	if grid[startY][startX] == wall {
		return -1
	}

	visited := make(map[visitedKey]bool)

	queue := []queueValue{
		{y: startY, x: startX, distance: 0},
	}

	for len(queue) > 0 {
		//nolint:varnamelen
		y := queue[0].y
		//nolint:varnamelen
		x := queue[0].x
		distance := queue[0].distance
		queue = queue[1:]

		if grid[y][x] == carrot {
			return distance
		}

		key := visitedKey{y: y, x: x}

		visited[key] = true

		for _, neighborOffset := range offsets() {
			neighborY := y + neighborOffset[0]
			neighborX := x + neighborOffset[1]
			neighborDistance := distance + 1

			if (neighborY >= 0) && (neighborY < len(grid)) && (neighborX >= 0) && (neighborX < len(grid[neighborY])) {
				// Helpful: Premature return
				if grid[neighborY][neighborX] == carrot {
					return neighborDistance
				}

				neighborKey := visitedKey{y: neighborY, x: neighborX}

				if !visited[neighborKey] && (grid[neighborY][neighborX] != wall) {
					queue = append(queue, queueValue{y: neighborY, x: neighborX, distance: neighborDistance})
				}
			}
		}
	}

	return -1
}
