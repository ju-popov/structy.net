package tribonacci

func Recursive(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 0
	}

	//nolint:gomnd
	if n == 2 {
		return 1
	}

	//nolint:gomnd
	return Recursive(n-1) + Recursive(n-2) + Recursive(n-3)
}
