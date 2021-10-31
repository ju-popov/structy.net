package shortestpath

func createGraph(edges [][2]string) map[string][]string {
	graph := make(map[string][]string)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

type nodeInfo struct {
	name     string
	distance int
}

func BreadthFirst(edges [][2]string, nodeA string, nodeB string) int {
	graph := createGraph(edges)
	visited := make(map[string]bool)

	queue := []nodeInfo{
		{
			name:     nodeA,
			distance: 0,
		},
	}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.name == nodeB {
			return node.distance
		}

		visited[node.name] = true

		for _, next := range graph[node.name] {
			if !visited[next] {
				queue = append(queue, nodeInfo{name: next, distance: node.distance + 1})
			}
		}
	}

	return -1
}
