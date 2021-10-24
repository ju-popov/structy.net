package intersection

func Iterative(a []int64, b []int64) []int64 {
	result := []int64{}

	set := map[int64]struct{}{}

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
