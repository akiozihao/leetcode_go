package daily

import "strconv"

func countAndSay(n int) string {
	res := "1"
	for i := 0; i < n-1; i++ {
		j, tmp := 0, ""
		for idx, c := range res {
			if c != rune(res[j]) {
				tmp += strconv.Itoa(idx-j) + string(res[j])
				j = idx
			}
		}
		res = tmp + strconv.Itoa(len(res)-j) + string(res[j])
	}
	return res
}
