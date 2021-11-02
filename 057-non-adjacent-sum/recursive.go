package nonadjacentsum

func helper(nums []int, pos int, memory []int) int {
	// no elements
	if pos >= len(nums) {
		return 0
	}

	// last element
	if pos == len(nums)-1 {
		return nums[pos]
	}

	if memory[pos] >= 0 {
		return memory[pos]
	}

	//nolint:gomnd
	withCurrent := nums[pos] + helper(nums, pos+2, memory)
	withoutCurrent := helper(nums, pos+1, memory)

	if withCurrent > withoutCurrent {
		memory[pos] = withCurrent
	} else {
		memory[pos] = withoutCurrent
	}

	return memory[pos]
}

func Recursive(nums []int) int {
	memory := make([]int, len((nums)))
	for i := range nums {
		memory[i] = -1
	}

	return helper(nums, 0, memory)
}
