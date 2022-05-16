package sanye

func findMaxLength(nums []int) (res int) {
	mp := map[int]int{0: -1}
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, has := mp[counter]; has {
			res = max(res, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
