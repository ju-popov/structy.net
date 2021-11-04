package semestersrequired

func createGraph(numCourses int, prereqs [][2]int) map[int][]int {
	// Create courses graph
	graph := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		graph[i] = make([]int, 0)
	}

	// Add prereqs
	for _, prereq := range prereqs {
		graph[prereq[0]] = append(graph[prereq[0]], prereq[1])
	}

	return graph
}

func explore(graph map[int][]int, course int, visited map[int]int) int {
	if value, ok := visited[course]; ok {
		return value
	}

	maxResult := 0

	for _, prereq := range graph[course] {
		if result := explore(graph, prereq, visited); result > maxResult {
			maxResult = result
		}
	}

	visited[course] = maxResult + 1

	return visited[course]
}

func DepthFirstRecursive(numCourses int, prereqs [][2]int) int {
	graph := createGraph(numCourses, prereqs)

	// Identify no prereqs courses
	visited := make(map[int]int)

	for course, prereqs := range graph {
		if len(prereqs) == 0 {
			visited[course] = 1
		}
	}

	maxResult := 0

	// Explore courses
	for course := range graph {
		if result := explore(graph, course, visited); result > maxResult {
			maxResult = result
		}
	}

	return maxResult
}
