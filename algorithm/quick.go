package algorithm

import "fmt"

func partition(list []int, left int, right int) int {
	pivot := list[left]

	for left < right {
		for left < right && list[right] >= pivot {
			right--
		}
		list[left] = list[right]

		for left < right && list[left] <= pivot {
			left++
		}
		list[right] = list[left]
	}

	list[left] = pivot

	fmt.Println(list)

	return left
}

func Quicksort(list []int, left int, right int) {
	if left < right {
		p := partition(list, left, right)
		Quicksort(list, left, p-1)
		Quicksort(list, p+1, right)
	}
}
