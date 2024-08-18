package algorithm

func Binary(arr []int, val int) int {
	if len(arr) == 0 {
		return -1
	}

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2

		if arr[mid] == val {
			return arr[mid]
		} else if arr[mid] < val {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
