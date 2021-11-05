package permutations

func Recursive(items []interface{}) [][]interface{} {
	if len(items) == 0 {
		return [][]interface{}{{}}
	}

	results := make([][]interface{}, 0)

	permutations := Recursive(items[1:])

	for _, permutation := range permutations {
		for index := 0; index <= len(permutation); index++ {
			result := make([]interface{}, 0)
			result = append(result, permutation[:index]...)
			result = append(result, items[0])
			result = append(result, permutation[index:]...)
			results = append(results, result)
		}
	}

	return results
}
