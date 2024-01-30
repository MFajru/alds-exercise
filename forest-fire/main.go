package main

import "fmt"

type node struct {
	value []int
	next  *node
}

type queue struct {
	head *node
	tail *node
	size int
}

func (Q *queue) dequeue() []int {
	if Q.head != nil {
		value := Q.head.value
		Q.head = Q.head.next
		Q.size--
		return value
	}
	return nil
}

func (Q *queue) enqueue(newValue []int) {
	newNode := &node{value: newValue, next: nil}
	if Q.tail != nil {
		fmt.Println("ok5")
		fmt.Println(newNode)
		fmt.Printf("Q tail: %d\n", Q.tail.value)
		Q.tail.next = newNode
		Q.tail = newNode
		Q.size++

		// fmt.Printf("Q head: %d", Q.head.value)
		fmt.Printf("Q tail: %d\n", Q.tail.value)
	} else {
		Q.tail = newNode
		Q.head = Q.tail
		Q.size++
	}
	fmt.Printf("Q head: %d", Q.head.value)
}

func isValid(lengthRow int, lengthCol int, vis [][]bool, row int, col int) bool {
	if row < 0 || col < 0 || row >= lengthRow || col >= lengthCol {
		fmt.Println("ok")
		return false
	}

	if vis[row][col] {
		fmt.Println("ok2")
		return false
	}
	fmt.Println("ok3")
	return true
}

func fireBFS(input [][]string, vis [][]bool, row int, col int) int {
	//Move directions vectors
	moveRow := []int{-1, 0, 1, 0}
	moveCol := []int{0, -1, 0, 1}

	lengthRow, lengthCol := len(input), len(input[0])

	var Q queue
	var countDay int

	//add starting point to visited
	Q.enqueue([]int{row, col})
	vis[row][col] = true

	fmt.Println(Q.head)
	fmt.Println(Q.tail)
	for Q.size > 0 {
		cell := Q.dequeue()
		fmt.Println(cell)
		x, y := cell[0], cell[1]
		fmt.Println(x, y)

		for i := 0; i < len(moveRow); i++ {
			fmt.Printf("x: %d, y: %d\n", x, y)
			adjX := x + moveRow[i]
			adjY := y + moveCol[i]
			fmt.Printf("adjX: %d adjY: %d\n", adjX, adjY)
			if isValid(lengthRow, lengthCol, vis, adjX, adjY) {
				Q.enqueue([]int{adjX, adjY})
				vis[adjX][adjY] = true
				break
			}
		}
		countDay++
	}
	return countDay

}

func findStartingPoint(input [][]string) (row int, col int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if input[i][j] == "X" {
				row = i
				col = j
				return row, col
			}
		}
	}
	return
}

func fillVisited(input [][]string) [][]bool {
	var visited [][]bool
	lengthRow, lengthCol := len(input), len(input[0])
	for i := 0; i < lengthRow; i++ {
		visited = append(visited, []bool{false})
		for j := 1; j < lengthCol; j++ {
			visited[i] = append(visited[i], false)
		}
	}
	return visited

}

func main() {
	input := [][]string{
		{"-", "-", "-"},
		{"-", "X", "-"},
		{"-", "-", "O"},
	}

	visited := fillVisited(input)

	row, col := findStartingPoint(input)

	fmt.Println(fireBFS(input, visited, row, col))

}
