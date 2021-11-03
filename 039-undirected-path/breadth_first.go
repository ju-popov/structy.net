package undirectedpath

func createGraphBreadthFirst(edges [][2]string) map[string][]string {
	graph := make(map[string][]string)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

func BreadthFirst(edges [][2]string, nodeA string, nodeB string) bool {
	graph := createGraphBreadthFirst(edges)

	visited := map[string]bool{}
	queue := []string{nodeA}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nodeB {
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
