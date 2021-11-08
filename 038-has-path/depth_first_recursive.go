package haspath

func DepthFirstRecursive(graph map[string][]string, src string, dst string) bool {
	if src == dst {
		return true
	}

	for _, neighbor := range graph[src] {
		if DepthFirstRecursive(graph, neighbor, dst) {
			return true
		}
	}

	return false
}
