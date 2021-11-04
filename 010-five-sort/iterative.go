package fivesort

const targetValue = 5

func Iterative(nums []int) []int {
	indexLeft := 0
	indexRight := len(nums) - 1

	for indexLeft < indexRight {
		switch {
		case nums[indexRight] == targetValue:
			indexRight--
		case nums[indexLeft] == targetValue:
			nums[indexLeft], nums[indexRight] = nums[indexRight], nums[indexLeft]
			indexLeft++
			indexRight--
		default:
			indexLeft++
		}
	}

	return nums
}
