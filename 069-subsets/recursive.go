package subsets

func helper(elements []string, results [][]string) [][]string {
	if len(elements) == 0 {
		result := []string{}
		results = append(results, result)

		return results
	}

	results = helper(elements[1:], results)

	for _, value := range results {
		result := append([]string{elements[0]}, value...)
		results = append(results, result)
	}

	return results
}

func Recursive(elements []string) [][]string {
	return helper(elements, make([][]string, 0))
}
