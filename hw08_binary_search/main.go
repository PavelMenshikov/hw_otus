package main

import "fmt"

func binarySearch(arr []int, target int) int {
	low, high := 0, len(arr)-1
	for low <= high {
		mid := low + (high-low)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	target := 11
	index := binarySearch(arr, target)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Printf("Элемент %d не найден\n", target)
	}
}
