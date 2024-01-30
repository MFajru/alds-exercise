package main

import "fmt"

func hasDuplicatesBruteForce(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				return true
			}
		}
	}
	return false
}

func hasDuplicatesMap(arr []int) bool {
	arrMap := make(map[int]int, 0)
	for _, num := range arr {
		arrMap[num]++
		if arrMap[num] > 1 {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 3, 5, 7, 10, 9}
	fmt.Println("Has Duplicates (Brute Force):", hasDuplicatesMap(arr))
}
