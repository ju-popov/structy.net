package haspath

func Recursive(graph map[string][]string, src string, dst string) bool {
	visited := make(map[string]bool)

	return helper(graph, src, dst, visited)
}

func helper(graph map[string][]string, src string, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}

	visited[src] = true

	for _, node := range graph[src] {
		if !visited[node] {
			if helper(graph, node, dst, visited) {
				return true
			}
		}
	}

	return false
}
