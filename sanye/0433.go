package sanye

func minMutation(start, end string, bank []string) int {
	if start == end {
		return 0
	}
	backSet := map[string]struct{}{}
	for _, s := range bank {
		backSet[s] = struct{}{}
	}
	if _, ok := backSet[end]; !ok {
		return -1
	}
	q := []string{start}
	for step := 0; q != nil; step++ {
		tmp := q
		q = nil
		for _, cur := range tmp {
			for i, x := range cur {
				for _, y := range "ACGT" {
					if y != x {
						nxt := cur[:i] + string(y) + cur[i+1:]
						if _, ok := backSet[nxt]; ok {
							if nxt == end {
								return step + 1
							}
							delete(backSet, nxt)
							q = append(q, nxt)
						}
					}
				}
			}
		}
	}
	return -1
}
