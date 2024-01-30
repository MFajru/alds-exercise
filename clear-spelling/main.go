package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertToNato(input string) string {
	natoAlphabet := map[string]string{
		"A": "Alfa",
		"B": "Bravo",
		"C": "Charlie",
		"D": "Delta",
		"E": "Echo",
		"F": "Foxtrot",
		"G": "Golf",
		"H": "Hotel",
		"I": "India",
		"J": "Juliett",
		"K": "Kilo",
		"L": "Lima",
		"M": "Mike",
		"N": "November",
		"O": "Oscar",
		"P": "Papa",
		"Q": "Quebec",
		"R": "Romeo",
		"S": "Sierra",
		"T": "Tango",
		"U": "Uniform",
		"V": "Victor",
		"W": "Whiskey",
		"X": "X-ray",
		"Y": "Yankee",
		"Z": "Zulu",
		" ": "Space",
	}

	modInput := strings.ToUpper(input)
	var natoString string

	for _, char := range modInput {
		natoString += fmt.Sprintf("%s ", natoAlphabet[string(char)])
	}
	return natoString
}

func main() {
	var input string
	fmt.Print("Input: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		fmt.Printf("Invalid input, %s. Please input string/words.\n", err)
	}

	result := convertToNato(input)
	fmt.Printf("Coverted to NATO phonetic alphabet: %s\n", result)
}
