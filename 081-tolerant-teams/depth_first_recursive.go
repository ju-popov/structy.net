package tolerantteams

func createGraph(rivalries [][2]string) map[string][]string {
	graph := make(map[string][]string)

	for _, rivalry := range rivalries {
		if _, ok := graph[rivalry[0]]; !ok {
			graph[rivalry[0]] = make([]string, 0)
		}

		if _, ok := graph[rivalry[1]]; !ok {
			graph[rivalry[1]] = make([]string, 0)
		}

		graph[rivalry[0]] = append(graph[rivalry[0]], rivalry[1])
		graph[rivalry[1]] = append(graph[rivalry[1]], rivalry[0])
	}

	return graph
}

func helper(graph map[string][]string, person string, team bool, assigned map[string]bool) bool {
	if value, ok := assigned[person]; ok {
		return value == team
	}

	assigned[person] = team

	nextTeam := !team

	for _, rivalry := range graph[person] {
		if !helper(graph, rivalry, nextTeam, assigned) {
			return false
		}
	}

	return true
}

func DepthFirstRecursive(rivalries [][2]string) bool {
	graph := createGraph(rivalries)
	assigned := make(map[string]bool)

	for person := range graph {
		if _, ok := assigned[person]; !ok {
			if !helper(graph, person, false, assigned) {
				return false
			}
		}
	}

	return true
}
