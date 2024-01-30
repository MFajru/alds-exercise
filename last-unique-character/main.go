package main

import "fmt"

func countCharMap(input string) map[string]int {
	charMap := map[string]int{}
	for _, char := range input {
		charMap[string(char)]++
	}
	return charMap
}

func findLastUniqueCharacters(input string, charMap map[string]int) string {
	for i := len(input) - 1; i >= 0; i-- {
		strInput := string(input[i])
		if charMap[strInput] == 1 {
			return strInput
		}
	}
	return ""

}

func main() {
	input := "ajjkkrrtt"
	charMap := countCharMap(input)
	fmt.Println(findLastUniqueCharacters(input, charMap))
}
