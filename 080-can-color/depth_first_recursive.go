package cancolor

func helper(graph map[string][]string, node string, currentColor bool, colored map[string]bool) bool {
	if value, ok := colored[node]; ok {
		return value == currentColor
	}

	colored[node] = currentColor

	neighbourColor := !currentColor

	for _, neighbour := range graph[node] {
		if !helper(graph, neighbour, neighbourColor, colored) {
			return false
		}
	}

	return true
}

func DepthFirstRecursive(graph map[string][]string) bool {
	colored := make(map[string]bool)

	for node := range graph {
		if _, ok := colored[node]; !ok {
			if !helper(graph, node, false, colored) {
				return false
			}
		}
	}

	return true
}
