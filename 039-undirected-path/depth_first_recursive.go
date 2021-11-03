package undirectedpath

func createGraphDepthFirstRecursive(edges [][2]string) map[string][]string {
	graph := make(map[string][]string)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

func helper(graph map[string][]string, src string, dst string, visited map[string]bool) bool {
	if src == dst {
		return true
	}

	visited[src] = true

	for _, neighbor := range graph[src] {
		if !visited[neighbor] {
			if helper(graph, neighbor, dst, visited) {
				return true
			}
		}
	}

	return false
}

func DepthFirstRecursive(edges [][2]string, nodeA string, nodeB string) bool {
	graph := createGraphDepthFirstRecursive(edges)
	visited := make(map[string]bool)

	return helper(graph, nodeA, nodeB, visited)
}
