package fib

//nolint:varnamelen
func helper(n int, memory map[int]int) int {
	if value, ok := memory[n]; ok {
		return value
	}

	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	//nolint:gomnd
	memory[n] = helper(n-1, memory) + helper(n-2, memory)

	return memory[n]
}

func Recursive(n int) int {
	memory := make(map[int]int)

	return helper(n, memory)
}
