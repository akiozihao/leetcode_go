package sanye

func isCovered(ranges [][]int, left int, right int) bool {
	for i := left; i <= right; i++ {
		ok := false
		for m := range ranges {
			if i >= ranges[m][0] && i <= ranges[m][1] {
				ok = true
				break
			}
		}
		if !ok {
			return false
		}
	}
	return true
}
