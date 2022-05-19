package sanye

func xorQueries(arr []int, queries [][]int) []int {
	n, m := len(arr), len(queries)
	prexor := make([]int, n+1)
	prexor[0] = 0
	for i, v := range arr {
		prexor[i+1] = prexor[i] ^ v
	}
	var helper func(prexor, query []int) int
	helper = func(prexor, query []int) int {
		return prexor[query[1]+1] ^ prexor[query[0]]
	}

	res := make([]int, m)
	for i := range queries {
		res[i] = helper(prexor, queries[i])
	}

	return res
}
