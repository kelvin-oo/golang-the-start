package main

import (
	"fmt"
)

var balance float64 = 5000

func main() {
	fmt.Println("welcome to go bank")
	fmt.Println("what would you like to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. withdraw money")
	fmt.Println("3. deposit money")
	fmt.Println("4. exit")

	option := operation()
	if option == 1 {
		fmt.Println("Your Balance is: ", balance)
		main()
	}
	if option == 4 {
		return
	}

	if option == 2 {
		newAmount := withdraw()
		fmt.Print("Your new balalnce is", newAmount)
		main()
	}
	if option == 3 {
		newAmount := add()
		fmt.Print("Your new balalnce is", newAmount)
		main()
	}
}

func operation() float64 {
	var option float64
	if _, err := fmt.Scan(&option); err != nil || option > 4 || option < 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		operation()
	}
	return option
}

func withdraw() float64 {
	var amount float64
	if _, err := fmt.Scan(&amount); err != nil || amount < 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		withdraw()
	}
	var newAmount float64 = balance - amount
	return newAmount
}

func add() float64 {
	var amount float64
	if _, err := fmt.Scan(&amount); err != nil || amount < 0 {
		fmt.Println("Invalid input. Please enter a positive number.")
		add()
	}
	var newAmount float64 = balance + amount
	return newAmount
}
