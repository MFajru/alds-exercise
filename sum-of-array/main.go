package main

import "fmt"

func recursiveSum(arr []int, n int) int {
	if n == 0 {
		return 0
	}
	return arr[n-1] + recursiveSum(arr, n-1)
}

func sumOfArray(arr []int, n int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	fmt.Println("Sum of elements (Recursive Sum):", sumOfArray(arr, len(arr)))
}
