package sanye

import "sort"

func findRightInterval(intervals [][]int) []int {
	for i := range intervals {
		intervals[i] = append(intervals[i], i)
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	n := len(intervals)
	res := make([]int, n)
	for _, p := range intervals {
		i := sort.Search(n, func(i int) bool { return intervals[i][0] >= p[1] })
		if i < n {
			res[p[2]] = intervals[i][2]
		} else {
			res[p[2]] = -1
		}
	}
	return res
}
