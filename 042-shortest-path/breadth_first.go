package shortestpath

type queueValue struct {
	node     string
	distance int
}

func createGraph(edges [][2]string) map[string][]string {
	graph := make(map[string][]string)

	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}

	return graph
}

func BreadthFirst(edges [][2]string, nodeA string, nodeB string) int {
	graph := createGraph(edges)
	visited := make(map[string]bool)

	queue := []queueValue{
		{
			node:     nodeA,
			distance: 0,
		},
	}

	for len(queue) > 0 {
		node := queue[0].node
		distance := queue[0].distance
		queue = queue[1:]

		if node == nodeB {
			return distance
		}

		visited[node] = true

		for _, neighbor := range graph[node] {
			if !visited[neighbor] {
				queue = append(queue, queueValue{node: neighbor, distance: distance + 1})
			}
		}
	}

	return -1
}
