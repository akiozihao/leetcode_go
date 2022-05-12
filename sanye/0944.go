package sanye

func minDeletionSize(strs []string) (res int) {
	for j := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if strs[i-1][j] > strs[i][j] {
				res++
				break
			}
		}
	}
	return
}
