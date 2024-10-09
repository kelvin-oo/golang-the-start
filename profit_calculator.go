package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter your revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter your tax rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTask := revenue - expenses
	earningsAfterTask := (earningsBeforeTask) - (taxRate / 100 * earningsBeforeTask)
	earningRatio := earningsBeforeTask / earningsAfterTask

	fmt.Println("Your earning before task is: ", earningsBeforeTask)
	fmt.Println("Your earning after task is: ", earningsAfterTask)
	fmt.Println("Your earning ratio is: ", earningRatio)
}
