package sanye

func diStringMatch(s string) []int {
	n := len(s)
	res := make([]int, n+1)
	low, high := 0, n
	for i, v := range s {
		if v == 'D' {
			res[i] = high
			high -= 1
		} else {
			res[i] = low
			low += 1
		}
	}
	res[n] = low
	return res
}
