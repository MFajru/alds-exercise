package main

import (
	"fmt"
	"strings"
)

func isVowels(char string) bool {
	return strings.Contains("aiueoAIUEO", char)
}

func stringToSlice(input string) []string {
	var sliceInput []string
	for _, char := range input {
		sliceInput = append(sliceInput, string(char))
	}
	return sliceInput
}

func reverseNonVowels(input string) []string {
	sliceInput := stringToSlice(input)
	low := 0
	high := len(sliceInput) - 1

	for low <= high {
		if isVowels(sliceInput[low]) && isVowels(sliceInput[high]) {
			low++
			high--
		} else if !isVowels(sliceInput[low]) && isVowels(sliceInput[high]) {
			high--
		} else if isVowels(sliceInput[low]) && !isVowels(sliceInput[high]) {
			low++
		} else {
			sliceInput[low], sliceInput[high] = sliceInput[high], sliceInput[low]
			low++
			high--
		}
	}

	return sliceInput
}

func main() {
	input := "tartar"

	var result string

	sliceResult := reverseNonVowels(input)

	for _, char := range sliceResult {
		result += char
	}
	fmt.Printf("Input: %s\n", input)
	fmt.Printf("Output: %s\n", result)

}
