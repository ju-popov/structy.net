package arraystepper

func helper(nums []int, index int, memory map[int]bool) bool {
	if index == len(nums)-1 {
		return true
	}

	if value, ok := memory[index]; ok {
		return value
	}

	result := false

	for i := 1; (i <= nums[index]) && !result; i++ {
		result = helper(nums, index+i, memory)
	}

	memory[index] = result

	return memory[index]
}

func Recursive(nums []int) bool {
	memory := make(map[int]bool)

	return helper(nums, 0, memory)
}
