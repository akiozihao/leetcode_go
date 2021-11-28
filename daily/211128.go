package daily


func findAnagrams(s string, p string)(res []int) {
	sLen,pLen := len(s),len(p)
	if sLen < pLen {
		return 
	}

	var sCount,pCount [26]int
	for i,ch := range p {
		sCount[s[i]-'a']++
		pCount[ch-'a']++
	}
	if sCount == pCount {
		res = append(res, 0)
	}
	for i:= pLen;i < sLen;i++ {
		sCount[s[i-pLen]-'a']--
		sCount[s[i]-'a']++
		if sCount == pCount {
			res = append(res, i-pLen + 1)
		}
	}
	return 
}