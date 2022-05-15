package sanye

func sumOddLengthSubarrays(arr []int) int {
	n := len(arr)
	preSum := make([]int, n+1)
	count := 0
	for i := 1; i <= n; i++ {
		count += arr[i-1]
		preSum[i] = count
	}
	res := 0
	for i := 1; i <= n; i += 2 {
		for j := 0; j < n-i; j++ {
			res += preSum[j+i] - preSum[j]
		}
	}
	return res
}
