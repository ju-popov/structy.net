package pairedparentheses

func Iterative(str string) bool {
	counter := 0

	for _, char := range str {
		switch char {
		case '(':
			counter++
		case ')':
			if counter == 0 {
				return false
			}
			counter--
		}
	}

	return counter == 0
}
