package sanye

func numRollsToTarget(n int, k int, target int) int {
	mod := int((1e9) + 7)
	if target < n || target > n*k {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, target+1)
	}
	for i := 1; i <= target; i++ {
		if i <= k {
			dp[1][i] = 1
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= target; j++ {
			for l := 1; l <= k; l++ {
				if j > l {
					dp[i][j] = (dp[i][j] + dp[i-1][j-l]) % mod
				}
			}
		}
	}
	return dp[n][target]
}
