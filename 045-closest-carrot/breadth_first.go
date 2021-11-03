package closestcarrot

type queueValue struct {
	row      int
	col      int
	distance int
}

type visitedKey struct {
	row int
	col int
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

func BreadthFirst(grid [][]string, startRow int, startCol int) int {
	if grid[startRow][startCol] == wall {
		return -1
	}

	visited := make(map[visitedKey]bool)

	queue := []queueValue{
		{row: startRow, col: startCol, distance: 0},
	}

	for len(queue) > 0 {
		row := queue[0].row
		col := queue[0].col
		distance := queue[0].distance
		queue = queue[1:]

		if grid[row][col] == carrot {
			return distance
		}

		key := visitedKey{row: row, col: col}

		visited[key] = true

		for _, neighborOffset := range offsets() {
			neighborRow := row + neighborOffset[0]
			neighborCol := col + neighborOffset[1]
			neighborDistance := distance + 1

			if (neighborRow >= 0) && (neighborRow < len(grid)) && (neighborCol >= 0) && (neighborCol < len(grid[neighborRow])) {
				// Helpful: Premature return
				if grid[neighborRow][neighborCol] == carrot {
					return neighborDistance
				}

				neighborKey := visitedKey{row: neighborRow, col: neighborCol}

				if !visited[neighborKey] && (grid[neighborRow][neighborCol] != wall) {
					queue = append(queue, queueValue{row: neighborRow, col: neighborCol, distance: neighborDistance})
				}
			}
		}
	}

	return -1
}
