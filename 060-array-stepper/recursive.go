package arraystepper

func helper(nums []int, index int, memory map[int]bool) bool {
	if value, ok := memory[index]; ok {
		return value
	}

	if index == len(nums)-1 {
		return true
	}

	for i := 1; i <= nums[index]; i++ {
		if helper(nums, index+i, memory) {
			memory[index] = true

			return memory[index]
		}
	}

	memory[index] = false

	return memory[index]
}

func Recursive(nums []int) bool {
	memory := make(map[int]bool)

	return helper(nums, 0, memory)
}
