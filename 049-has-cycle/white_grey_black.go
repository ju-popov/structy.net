package hascycle

func isCyclic(graph map[string][]string, node string, visiting map[string]bool, visited map[string]bool) bool {
	if visited[node] {
		return false
	}

	if visiting[node] {
		return true
	}

	visiting[node] = true

	for _, neighbor := range graph[node] {
		if isCyclic(graph, neighbor, visiting, visited) {
			return true
		}
	}

	visiting[node] = false
	visited[node] = true

	return false
}

func WhiteGreyBlack(graph map[string][]string) bool {
	visiting := make(map[string]bool)
	visited := make(map[string]bool)

	for node := range graph {
		if isCyclic(graph, node, visiting, visited) {
			return true
		}
	}

	return false
}
