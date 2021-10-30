package connectedcomponentscount

func explore(graph map[int][]int, node int, visited map[int]bool) bool {
	if visited[node] {
		return false
	}

	visited[node] = true

	for _, next := range graph[node] {
		if !visited[next] {
			explore(graph, next, visited)
		}
	}

	return true
}

func Recursive(graph map[int][]int) int {
	result := 0

	visited := map[int]bool{}

	for node := range graph {
		if explore(graph, node, visited) {
			result++
		}
	}

	return result
}
