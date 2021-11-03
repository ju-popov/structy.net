package connectedcomponentscount

func explore(graph map[int][]int, node int, visited map[int]bool) bool {
	if visited[node] {
		return false
	}

	visited[node] = true

	for _, neighbor := range graph[node] {
		if !visited[neighbor] {
			explore(graph, neighbor, visited)
		}
	}

	return true
}

func DepthFirstRecursive(graph map[int][]int) int {
	result := 0

	visited := map[int]bool{}

	for neighbor := range graph {
		if explore(graph, neighbor, visited) {
			result++
		}
	}

	return result
}
