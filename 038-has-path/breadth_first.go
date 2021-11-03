package haspath

func BreadthFirst(graph map[string][]string, src string, dst string) bool {
	visited := map[string]bool{}
	queue := []string{src}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == dst {
			return true
		}

		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, neighbor)
			}
		}
	}

	return false
}
