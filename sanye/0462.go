package sanye

import "sort"

func minMoves2(nums []int) int {
	res, i, j := 0, 0, len(nums)-1
	sort.Ints(nums)
	for i < j {
		res += nums[j] - nums[i]
		j--
		i++
	}
	return res
}
