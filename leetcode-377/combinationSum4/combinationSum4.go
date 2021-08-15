package combinationSum4

import (
	"sort"
)

func CombinationSum4(nums []int, target int) int {
	sort.Ints(nums)
	if nums[0] > target {
		return 0
	}
	if nums[0] == target {
		return 1
	}
	var tempResult int
	var resultCnt int
	for i := 0; i < len(nums); i++ {
		for tempResult < target {
			tempResult += nums[i]
		}
	}
	return resultCnt
}
