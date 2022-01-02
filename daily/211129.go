package daily

import "sort"

func kthSmallestPrimeFraction(arr []int, k int) []int {
	n := len(arr)
	type pair struct {
		x,y int
	}
	key := make([]pair,0,n*(n-1)/2)
	for i,x := range arr {
		for _,y := range arr[i+1:] {
			key = append(key, pair{x,y})
		}
	}
	sort.Slice(key,func(i,j int) bool{
		return key[i].x * key[j].y < key[i].y * key[j].x
	})
	return []int{key[k-1].x,key[k-1].y}
}