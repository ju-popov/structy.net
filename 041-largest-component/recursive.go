package largestcomponent

func explore(graph map[int][]int, node int, visited map[int]bool) int {
	if visited[node] {
		return 0
	}

	result := 1
	visited[node] = true

	for _, next := range graph[node] {
		if !visited[next] {
			result += explore(graph, next, visited)
		}
	}

	return result
}

func Recursive(graph map[int][]int) int {
	maxResult := 0

	visited := map[int]bool{}

	for node := range graph {
		if result := explore(graph, node, visited); result > maxResult {
			maxResult = result
		}
	}

	return maxResult
}
