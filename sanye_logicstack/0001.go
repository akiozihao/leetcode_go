func twoSum(nums []int, target int) []int {
	var res [2]int
	if len(nums) == 0 {
		return res
	}
	var m = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if val, ok := m[nums[i]]; ok {
			res[0] = m[nums[i]]
			res[1] = i
			return res
		}
		m[target-nums[i]] = i
	}
	return res
}