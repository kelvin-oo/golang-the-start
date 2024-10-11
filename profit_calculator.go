package main

// import "fmt"

// func main() {

// 	revenue, expenses, taxRate := getInputs()
// 	var ebt, eat, earningRatio = calculateValues(revenue, expenses, taxRate)

// 	fmt.Println("Your earning before task is: ", ebt)
// 	fmt.Println("Your earning after task is: ", eat)
// 	fmt.Println("Your earning ratio is: ", earningRatio)
// }

// func calculateValues(revenue, expenses, taxRate float64) (ebt, eat, earningRatio float64) {
// 	ebt = revenue - expenses
// 	eat = (ebt) - (taxRate / 100 * ebt)
// 	earningRatio = ebt / eat
// 	return ebt, eat, earningRatio
// }

// func getInputs() (float64, float64, float64) {
// 	var revenue, expenses, taxRate float64

// 	fmt.Print("Enter Revenue: $")
// 	if _, err := fmt.Scan(&revenue); err != nil || revenue <= 0 {
// 		fmt.Println("Invalid input. Please enter a positive number.")
// 		return getInputs()
// 	}

// 	fmt.Print("Enter expenses: ")
// 	if _, err := fmt.Scan(&expenses); err != nil || expenses < 0 {
// 		fmt.Println("Invalid input. Please en ter a non-negative number.")
// 		return getInputs()
// 	}

// 	fmt.Print("Enter tax rate: ")
// 	if _, err := fmt.Scan(&taxRate); err != nil || taxRate <= 0 {
// 		fmt.Println("Invalid input. Please enter a positive number.")
// 		return getInputs()
// 	}

// 	return revenue, expenses, taxRate
// }
