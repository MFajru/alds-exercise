package main

import "fmt"

func binarySearch(arr []int, target int, length int) int {
	low := 0
	high := length - 1

	for low <= high {
		mid := (low + high) / 2

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
	input := []int{2, 4, 6, 8, 10, 12, 14}
	length := len(input)
	target := 14
	search := binarySearch(input, target, length)
	fmt.Printf("%d is found in index %d\n", target, search)
}
