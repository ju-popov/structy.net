package maxvalue

func Simple(nums []float64) float64 {
	var result float64

	for i, num := range nums {
		if i == 0 {
			result = num
		} else if num > result {
			result = num
		}
	}

	return result
}
