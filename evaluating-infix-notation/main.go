package main

import (
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	value string
	next  *node
}

type stack struct {
	top    *node
	length int
}

func (S *stack) pop() string {
	if S.top != nil {
		popValue := S.top.value
		S.top = S.top.next
		S.length--
		return popValue
	}
	return ""
}

func (S *stack) push(newVal string) {
	newNode := &node{value: newVal, next: nil}
	if S.top == nil {
		S.top = newNode
		S.length++
	} else {
		newNode.next = S.top
		S.top = newNode
		S.length++
	}
}

func numOperation(operand1 string, operand2 string, operator string) string {
	intOperand1, _ := strconv.Atoi(operand1)
	intOperand2, _ := strconv.Atoi(operand2)

	var result int
	switch operator {
	case "*":
		result = intOperand1 * intOperand2
	case "/":
		result = intOperand1 / intOperand2
	case "+":
		result = intOperand1 + intOperand2
	case "-":
		result = intOperand1 - intOperand2
	}

	strResult := strconv.Itoa(result)

	return strResult
}

func main() {
	input := "2 * 2 + 7"
	sliceInput := strings.Split(input, " ")

	number := stack{}
	operator := stack{}

	for _, char := range sliceInput {
		if _, err := strconv.Atoi(char); err == nil {
			number.push(char)
		} else if char != "(" && char != ")" {
			for operator.top != nil && (operator.top.value == "*" || operator.top.value == "/") {
				operand2 := number.pop()
				operand1 := number.pop()
				op := operator.pop()
				number.push(numOperation(operand1, operand2, op))
			}
			operator.push(char)
		} else if char == ")" {
			for operator.top != nil && operator.top.value != "(" {
				operand2 := number.pop()
				operand1 := number.pop()
				op := operator.pop()
				number.push(numOperation(operand1, operand2, op))
			}
			operator.pop()
		} else {
			operator.push(char)
		}
	}

	for operator.top != nil {
		operand2 := number.pop()
		operand1 := number.pop()
		op := operator.pop()
		number.push(numOperation(operand1, operand2, op))
	}

	fmt.Println(number.pop())
}
