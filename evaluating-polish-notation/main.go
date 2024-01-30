package main

import (
	"fmt"
	"strconv"
	"strings"
)

// buat slice baru untuk menampung stack
// cek slice of string, punya value masukkan ke stack, selama belum menemukan operator
// berhenti setelah ketemu operator, operasikan value tersebut dengan operator yang ditemukan (pop value dari stack)
// hanya dua value yang bisa diambil dari stack
// masukkan hasil operasi ke stack baru bersama dengan sisa karakter
// lakukan berulang

func main() {
	// input := "+ 10 2"
	input := "/ 1 * 1 + 3 4"
	fmt.Printf("Input: %s\n", input)
	sliceInput := strings.Split(input, " ")
	var stack []int

	for i := len(sliceInput) - 1; i >= 0; i-- {
		if strings.Contains("*/+-", sliceInput[i]) {
			fmt.Println(stack)
			operand1 := stack[len(stack)-2]
			operand2 := stack[len(stack)-1]
			switch sliceInput[i] {
			case "*":
				tempResult := operand1 * operand2
				stack = append(stack, tempResult)
			case "/":
				tempResult := operand1 / operand2
				stack = append(stack, tempResult)
			case "+":
				tempResult := operand1 + operand2
				stack = append(stack, tempResult)
			case "-":
				tempResult := operand1 - operand2
				stack = append(stack, tempResult)
			}
		} else {
			intChar, _ := strconv.Atoi(sliceInput[i])
			stack = append(stack, intChar)
		}
	}
	fmt.Println(stack)
	fmt.Printf("Result: %d\n", stack[len(stack)-1])

}
