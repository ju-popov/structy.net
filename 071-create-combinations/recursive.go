package createcombinations

//nolint:varnamelen
func Recursive(items []interface{}, k int) [][]interface{} {
	results := make([][]interface{}, 0)

	if k > len(items) {
		return results
	}

	if k == 0 {
		result := make([]interface{}, 0)
		results = append(results, result)

		return results
	}

	combinations := Recursive(items[1:], k-1)
	for _, combination := range combinations {
		result := []interface{}{items[0]}
		result = append(result, combination...)
		results = append(results, result)
	}

	results = append(results, Recursive(items[1:], k)...)

	return results
}
