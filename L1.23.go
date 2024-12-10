package main

import (
	"fmt"
)

func removeAppend(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}
	return append(slice[:i], slice[i+1:]...)
}

func removeLoop(slice []int, i int) []int {
	if i < 0 || i >= len(slice) {
		return slice
	}

	result := make([]int, 0, len(slice)-1)
	for j, v := range slice {
		if j != i {
			result = append(result, v)
		}
	}
	return result
}

func removePlace(slice *[]int, i int) { //Этот метод изменяет оригинальный срез, перемещая элементы
	if i < 0 || i >= len(*slice) {
		return
	}
	*slice = append((*slice)[:i], (*slice)[i+1:]...)
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Исходный срез:", slice1)
	slice1 = removeAppend(slice1, 6)
	fmt.Println("После удаления элемента с индексом 6 (append):", slice1)

	slice2 := []int{10, 20, 30, 40, 50}
	fmt.Println("Исходный срез (loop):", slice2)
	slice2 = removeLoop(slice2, 1)
	fmt.Println("После удаления элемента с индексом 1 (loop):", slice2)

	slice3 := []int{100, 200, 300, 400, 500}
	fmt.Println("Исходный срез (in-place):", slice3)
	removePlace(&slice3, 2)
	fmt.Println("После удаления элемента с индексом 2 (in-place):", slice3)
	removePlace(&slice3, 0)
	fmt.Println("После удаления элемента с индексом 0 (in-place):", slice3)
	removePlace(&slice3, 1)
	fmt.Println("После удаления элемента с индексом 1 (in-place):", slice3)
}
