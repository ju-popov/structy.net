package fivesort

func find(nums []int64, minIndex int, maxIndex int, value int64) int {
	for index := minIndex; index <= maxIndex; index++ {
		if nums[index] == value {
			return index
		}
	}

	return -1
}

const targetValue = int64(5)

func Simple(nums []int64) []int64 {
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
