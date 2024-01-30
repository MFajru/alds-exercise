package main

import "fmt"

const maxCapacity = 5

func main() {
	input := []string{"newspaper", "cartoon", "fabel", "book", "tabloid"}
	next := []string{"catalog", "pen"}

	if len(next) > maxCapacity {
		next = next[len(next)-maxCapacity:]
	}

	popMaxIndex := len(input) - len(next)
	result := input[:popMaxIndex]
	result = append(result, next...)

	fmt.Println(result)
}
