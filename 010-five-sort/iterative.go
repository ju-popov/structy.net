package fivesort

func find(nums []int, minIndex int, maxIndex int, value int) int {
	for index := minIndex; index <= maxIndex; index++ {
		if nums[index] == value {
			return index
		}
	}

	return -1
}

const targetValue = int(5)

func Iterative(nums []int) []int {
	indexLeft := -1
	for indexRight := len(nums) - 1; indexLeft < indexRight; indexRight-- {
		if nums[indexRight] != targetValue {
			if targetIndex := find(nums, indexLeft+1, indexRight-1, targetValue); targetIndex != -1 {
				nums[targetIndex] = nums[indexRight]
				nums[indexRight] = targetValue
				indexLeft = targetIndex
			} else {
				return nums
			}
		}
	}

	return nums
}
