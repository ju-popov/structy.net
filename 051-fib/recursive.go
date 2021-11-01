package fib

func Recursive(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	//nolint:gomnd
	return Recursive(n-1) + Recursive(n-2)
}
