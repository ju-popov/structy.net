package prereqspossible

func createGraph(numCourses int, prereqs [][2]int) map[int][]int {
	graph := make(map[int][]int)

	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	for _, prereq := range prereqs {
		graph[prereq[0]] = append(graph[prereq[0]], prereq[1])
	}

	return graph
}

func isCyclic(graph map[int][]int, node int, visiting map[int]bool, visited map[int]bool) bool {
	if visited[node] {
		return false
	}

	if visiting[node] {
		return true
	}

	visiting[node] = true

	for _, neighbor := range graph[node] {
		if isCyclic(graph, neighbor, visiting, visited) {
			return true
		}
	}

	visiting[node] = false
	visited[node] = true

	return false
}

func WhiteGreyBlack(numCourses int, prereqs [][2]int) bool {
	graph := createGraph(numCourses, prereqs)

	visiting := make(map[int]bool)
	visited := make(map[int]bool)

	for node := range graph {
		if isCyclic(graph, node, visiting, visited) {
			return false
		}
	}

	return true
}
