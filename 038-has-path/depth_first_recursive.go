package haspath

func helper(graph map[string][]string, src string, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}

	visited[src] = true

	for _, next := range graph[src] {
		if !visited[next] {
			if helper(graph, next, dst, visited) {
				return true
			}
		}
	}

	return false
}

func DepthFirstRecursive(graph map[string][]string, src string, dst string) bool {
	visited := make(map[string]bool)

	return helper(graph, src, dst, visited)
}
