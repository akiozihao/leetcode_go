package sanye

import (
	"math"
	"sort"
)

func numSubarrayProductLessThanK(nums []int, k int) (ans int) {
	if k == 0 {
		return 0
	}
	n := len(nums)
	logPrefix := make([]float64, n+1)
	for i, num := range nums {
		logPrefix[i+1] = logPrefix[i] + math.Log(float64(num))
	}
	logK := math.Log(float64(k))
	for j := 1; j <= n; j++ {
		l := sort.SearchFloat64s(logPrefix[:j], logPrefix[j]-logK+1e-10)
		ans += j - l
	}
	return 
}
