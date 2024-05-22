package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func BinarySearch[T constraints.Ordered](arr []T, target T) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2 // чтобы избежать переполнения

		if arr[mid] == target {
			return mid
		}

		if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // элемент не найден
}

func main() {
	myList := []int{1, 3, 5, 7, 9, 13, 15, 17, 19, 20, 25, 34, 44, 45, 55, 56, 67, 78, 90, 100}
	myFloatList := []float32{1.1, 1.6, 2.5, 4.5, 6.1, 7.6, 11.4, 15.5}
	myStringList := []string{"ab", "ac", "b", "bcd", "fav", "fbd"}

	fmt.Println(BinarySearch(myList, 9))
	fmt.Println(BinarySearch(myList, -1))

	fmt.Println(BinarySearch(myList, 7))
	fmt.Println(BinarySearch(myFloatList, 6.1))
	fmt.Println(BinarySearch(myStringList, "fav"))
}
