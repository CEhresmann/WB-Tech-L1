package main

import (
	"fmt"
)

func bsIterative(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func bsRecursive(arr []int, target int, left int, right int) int {
	if left > right {
		return -1
	}
	mid := left + (right-left)/2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		return bsRecursive(arr, target, mid+1, right)
	} else {
		return bsRecursive(arr, target, left, mid-1)
	}
}

func bsWSlice(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	mid := len(arr) / 2
	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		index := bsWSlice(arr[mid+1:], target)
		if index == -1 {
			return -1
		}
		return mid + 1 + index
	} else {
		return bsWSlice(arr[:mid], target)
	}
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 7

	index1 := bsIterative(arr, target)
	fmt.Printf("Итеративный поиск: Элемент %d найден на индексе %d.\n", target, index1)

	index2 := bsRecursive(arr, target, 0, len(arr)-1)
	fmt.Printf("Рекурсивный поиск: Элемент %d найден на индексе %d.\n", target, index2)

	index3 := bsWSlice(arr, target)
	fmt.Printf("Поиск с использованием срезов: Элемент %d найден на индексе %d.\n", target, index3)
}
