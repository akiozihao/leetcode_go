package sanye

func lastStoneWeightII(stones []int) int {
	if len(stones) == 1 {
		return stones[0]
	}
	sum := 0
	for _, v := range stones {
		sum += v
	}
	n, m := len(stones), sum/2
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, m+1)
	}
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	dp[0][stones[0]] = true

	for i := 1; i < n; i++ {
		v := stones[i]
		for j := 1; j <= m; j++ {
			if j < v {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-v]
			}
		}
	}
	for j := m; j >= 0; j-- {
		if dp[n-1][j] {
			return sum - 2*j
		}
	}
	return -1
}
