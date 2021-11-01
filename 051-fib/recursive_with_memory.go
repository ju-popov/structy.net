package fib

func helperWithMemory(n int, memory map[int]int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	if value, ok := memory[n]; ok {
		return value
	}

	//nolint:gomnd
	memory[n] = helperWithMemory(n-1, memory) + helperWithMemory(n-2, memory)

	return memory[n]
}

func RecursiveWithMemory(n int) int {
	memory := make(map[int]int)

	return helperWithMemory(n, memory)
}
