package intersection

//nolint:varnamelen
func Iterative(a []int, b []int) []int {
	result := make([]int, 0)

	set := make(map[int]struct{})

	for _, element := range a {
		set[element] = struct{}{}
	}

	for _, element := range b {
		if _, ok := set[element]; ok {
			result = append(result, element)
		}
	}

	return result
}
