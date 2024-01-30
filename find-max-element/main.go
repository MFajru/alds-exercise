package main

import "fmt"

func findMaxElement(arr []int) (int, error) {
	if len(arr) == 0 {
		return 0, fmt.Errorf("array is empty")
	}

	maxElement := arr[0]
	fmt.Println(arr[0])
	for i := 1; i < len(arr); i++ {
		fmt.Println(arr[i])
		if arr[i] > maxElement {
			maxElement = arr[i]
		}
	}
	return maxElement, nil
}

func main() {
	arr := []int{}
	maxElement, err := findMaxElement(arr)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Maximum Element:", maxElement)
	}
}

//if array empty
