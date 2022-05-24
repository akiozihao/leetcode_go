package sanye

import "sort"

func findUnsortedSubarray(nums []int) int {
	if sort.IntsAreSorted(nums) {
		return 0
	}
	numsSorted := append([]int(nil), nums...)
	sort.Ints(numsSorted)
	l, r := 0, len(nums)-1
	for nums[l] == numsSorted[l] {
		l++
	}
	for nums[r] == numsSorted[r] {
		r--
	}
	return r - l + 1
}
