package nestingscore

func Iterative(str string) int {
	stack := []int{0}

	for _, char := range str {
		if char == '[' {
			stack = append(stack, 0)
		} else {
			value := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if value == 0 {
				stack[len(stack)-1]++
			} else {
				//nolint:gomnd
				stack[len(stack)-1] += 2 * value
			}
		}
	}

	return stack[0]
}
