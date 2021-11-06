package rarerouting

func createGraph(n int, roads [][2]int) map[int][]int {
	graph := make(map[int][]int)

	for city := 0; city < n; city++ {
		graph[0] = make([]int, 0)
	}

	for _, road := range roads {
		graph[road[0]] = append(graph[road[0]], road[1])
		graph[road[1]] = append(graph[road[1]], road[0])
	}

	return graph
}

func isCyclic(graph map[int][]int, city int, lastCity int, visited map[int]bool) bool {
	if visited[city] {
		return true
	}

	visited[city] = true

	for _, neighbour := range graph[city] {
		if neighbour != lastCity {
			if isCyclic(graph, neighbour, city, visited) {
				return true
			}
		}
	}

	return false
}

//nolint:varnamelen
func DepthFirstRecursive(n int, roads [][2]int) bool {
	//nolint:ifshort
	graph := createGraph(n, roads)
	//nolint:ifshort
	visited := make(map[int]bool)

	if isCyclic(graph, 0, -1, visited) {
		return false
	}

	for city := 1; city < n; city++ {
		if !visited[city] {
			return false
		}
	}

	return true
}
