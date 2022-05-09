package sanye

func findDuplicates(nums []int) (res []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		for nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}

	for i, num := range nums {
		if num-1 != i {
			res = append(res, num)
		}
	}
	return
}
