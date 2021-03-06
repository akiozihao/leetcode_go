package daily

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)-1
	var mid int
	for left < right {
		mid = left + (right-left)/2
		if arr[mid] < arr[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return right
}
