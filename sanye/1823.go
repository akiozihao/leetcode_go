package sanye

func findTheWinner(n int, k int) int {
	q := make([]int, n)
	for i := range q {
		q[i] = i + 1
	}
	for len(q) > 1 {
		for i := 1; i < k; i++ {
			q = append(q, q[0])[1:]
		}
		q = q[1:]
	}
	return q[0]
}
