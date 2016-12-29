package bsearch

func BSearch(array []int, el int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := low + (high-low)/2
		if array[mid] == el {
			return mid
		} else if el > array[mid] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
