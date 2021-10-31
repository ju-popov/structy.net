package closestcarrot

type cellData struct {
	row      int
	col      int
	distance int
}

const (
	wall   = "X"
	carrot = "C"
)

func BreadthFirst(grid [][]string, startRow int, startCol int) int {
	if grid[startRow][startCol] == wall {
		return -1
	}

	visited := make([][]bool, len(grid))
	for index, row := range grid {
		visited[index] = make([]bool, len(row))
	}

	queue := []cellData{
		{row: startRow, col: startCol, distance: 0},
	}

	for len(queue) > 0 {
		cell := queue[0]
		queue = queue[1:]

		if grid[cell.row][cell.col] == carrot {
			return cell.distance
		}

		visited[cell.row][cell.col] = true

		for _, next := range [4][2]int{
			{-1, 0}, // up
			{+1, 0}, // down
			{0, -1}, // left
			{0, +1}, // right
		} {
			nextRow, nextCol, nextDistance := cell.row+next[0], cell.col+next[1], cell.distance+1

			if (nextRow >= 0) && (nextRow < len(grid)) && (nextCol >= 0) && (nextCol < len(grid[nextRow])) {
				// Helpful: Premature return
				if grid[nextRow][nextCol] == carrot {
					return nextDistance
				}

				if !visited[nextRow][nextCol] && (grid[nextRow][nextCol] != wall) {
					queue = append(queue, cellData{row: nextRow, col: nextCol, distance: nextDistance})
				}
			}
		}
	}

	return -1
}
