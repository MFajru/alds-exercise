package main

import "fmt"

func dequeque(front *int, lengthNext int) {
	*front += lengthNext
}

func enqueue(result []string, next []string) []string {
	result = append(result, next...)
	return result
}

func queue(input []string, next []string) []string {
	front := 0
	ptrFront := &front
	result := input

	dequeque(ptrFront, len(next))
	result = enqueue(result, next)

	return result[front:]
}

func main() {
	input := []string{"Dani", "Rian", "Toni", "Loni", "Poki", "Vali", "Juli", "Kris", "Veny", "Luis"}
	next := []string{"Edward", "Same"}

	fmt.Println(input)
	fmt.Println(queue(input, next))

}
