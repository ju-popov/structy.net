package haspath

func BreadthFirst(graph map[string][]string, src string, dst string) bool {
	queue := []string{src}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == dst {
			return true
		}

		queue = append(queue, graph[node]...)
	}

	return false
}
