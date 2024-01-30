package main

import (
	"fmt"
	"strings"
)

func modifyString(input string) []string {
	modInput := strings.ToLower(strings.ReplaceAll(input, ".", ""))
	sliceString := strings.Split(modInput, " ")

	return sliceString
}

func countWord(sliceString []string) (map[string]int, []string) {
	var sortedKey []string
	wordsMap := map[string]int{}

	for _, word := range sliceString {
		_, exist := wordsMap[word]
		if exist {
			wordsMap[word]++
		} else {
			wordsMap[word] = 1
			sortedKey = append(sortedKey, word)
		}
	}

	return wordsMap, sortedKey
}

func main() {
	input := "The quick brown fox jumps over the lazy dog. The dog barks loudly."

	sliceString := modifyString(input)

	wordsMap, sortedKey := countWord(sliceString)

	for _, word := range sortedKey {
		fmt.Printf("%s: %d\n", word, wordsMap[word])
	}

}
