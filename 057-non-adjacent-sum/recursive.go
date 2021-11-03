package nonadjacentsum

func maxInt(values ...int) int {
	maxValue := values[0]

	for i := 1; i < len(values); i++ {
		if values[i] > maxValue {
			maxValue = values[i]
		}
	}

	return maxValue
}

func helper(nums []int, pos int, memory map[int]int) int {
	if value, ok := memory[pos]; ok {
		return value
	}

	// no elements
	if pos >= len(nums) {
		return 0
	}

	//nolint:gomnd
	withCurrent := nums[pos] + helper(nums, pos+2, memory)
	withoutCurrent := helper(nums, pos+1, memory)

	memory[pos] = maxInt(withCurrent, withoutCurrent)

	return memory[pos]
}

func Recursive(nums []int) int {
	memory := make(map[int]int)

	return helper(nums, 0, memory)
}
