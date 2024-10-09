package main

// import (
// 	"fmt"
// 	"math"
// )

// const inflationRate = 2.5

// func main() {
// 	investmentAmount, ratePerYear, years := getInputs()
// 	futureValue, inflationAdjustedValue := calculateValues(inflationRate, investmentAmount, ratePerYear, years)

// 	fmt.Printf("Future Value: $%.2f\n", futureValue)
// 	fmt.Printf("Inflation Adjusted Value: $%.2f\n", inflationAdjustedValue)
// }

// func getInputs() (float64, float64, float64) {
// 	var investmentAmount, ratePerYear, years float64

// 	fmt.Print("Enter investment amount: $")
// 	if _, err := fmt.Scan(&investmentAmount); err != nil || investmentAmount <= 0 {
// 		fmt.Println("Invalid input. Please enter a positive number.")
// 		return getInputs()
// 	}

// 	fmt.Print("Enter annual interest rate (%): ")
// 	if _, err := fmt.Scan(&ratePerYear); err != nil || ratePerYear < 0 {
// 		fmt.Println("Invalid input. Please enter a non-negative number.")
// 		return getInputs()
// 	}

// 	fmt.Print("Enter number of years: ")
// 	if _, err := fmt.Scan(&years); err != nil || years <= 0 {
// 		fmt.Println("Invalid input. Please enter a positive number.")
// 		return getInputs()
// 	}

// 	return investmentAmount, ratePerYear, years
// }

// func calculateValues(inflationRate, investmentAmount, ratePerYear, years float64) (futureValue, inflationAdjustedValue float64) {
// 	// Calculate future value
// 	futureValue = investmentAmount * math.Pow(1+ratePerYear/100, years)

// 	// Calculate inflation-adjusted value
// 	inflationAdjustedValue = futureValue / math.Pow(1+inflationRate/100, years)

// 	return futureValue, inflationAdjustedValue
// }
