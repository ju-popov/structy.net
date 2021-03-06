package longestpath

func explore(graph map[string][]string, node string, visited map[string]int) int {
	if value, ok := visited[node]; ok {
		return value
	}

	maxResult := 0

	for _, neighbor := range graph[node] {
		result := explore(graph, neighbor, visited)
		if result > maxResult {
			maxResult = result
		}
	}

	visited[node] = maxResult + 1

	return visited[node]
}

func DepthFirstRecursive(graph map[string][]string) int {
	visited := make(map[string]int)

	// Mark end-nodes
	for node, neighbors := range graph {
		if len(neighbors) == 0 {
			visited[node] = 0
		}
	}

	maxResult := 0

	for node := range graph {
		if result := explore(graph, node, visited); result > maxResult {
			maxResult = result
		}
	}

	return maxResult
}
