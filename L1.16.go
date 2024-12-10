package main

import (
	"fmt"
	"sort"
)

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	left := []int{}
	right := []int{}

	for _, v := range arr[1:] {
		if v < pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	return append(append(quicksort(left), pivot), quicksort(right)...)
}

func main() {
	arr := IntSlice{3, 6, 8, 10, 1, 2, 1, 345, 56, 3256, 99, 9, 546432, 11}
	sort.Sort(arr)
	fmt.Println("Отсортированный массив с использованием sort.Sort:", arr)

	sortedArr := quicksort(arr)
	fmt.Println("Отсортированный массив с рекурсивной функцией:", sortedArr)

}
