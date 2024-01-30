package main

import "fmt"

func main() {
	carType := map[string]int{
		"Small":  1000,
		"Medium": 2000,
		"Big":    3000,
	}

	var car string
	var balance int

	fmt.Print("Car Type: ")
	_, carErr := fmt.Scanln(&car)

	fmt.Print("E-money Balance: ")
	_, balanceErr := fmt.Scanln(&balance)

	_, exist := carType[car]
	if !exist || carErr != nil {
		fmt.Println("Invalid car type entered.")
		return
	}

	remainBalance := balance - carType[car]
	if balance < carType[car] || balanceErr != nil {
		fmt.Println("Insufficient balance.")
		return
	}

	fmt.Printf("\nCost: %d\n", carType[car])
	fmt.Printf("Car type: %s\n", car)
	fmt.Printf("Remaining balance: %d\n", remainBalance)

}
