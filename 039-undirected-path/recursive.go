package undirectedpath

func createGraphRecursive(edges [][2]string) map[string][]string {
	graph := make(map[string][]string)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

func Recursive(edges [][2]string, nodeA string, nodeB string) bool {
	visited := make(map[string]bool)
	graph := createGraphRecursive(edges)

	return helper(graph, nodeA, nodeB, visited)
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
