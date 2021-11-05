package befittingbrackets

func Iterative(str string) bool {
	stack := make([]int32, 0)

	brackets := map[int32]int32{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	for _, char := range str {
		//nolint:gocritic
		if value, ok := brackets[char]; ok {
			stack = append(stack, value)
		} else if (len(stack) == 0) || (stack[len(stack)-1] != char) {
			return false
		} else {
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
