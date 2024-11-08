package main

import "fmt"

func main() {
	// numbers := [5]float64{3.14, 2.71, 1.41, 1.61, 0.57}
	// numbers[3] = 22.5
	// fmt.Println(numbers[3])

	// lenght := len(numbers)
	// fmt.Println(lenght)

	// featuredNumbers := numbers[1:3]
	// fmt.Println(featuredNumbers)

	// someNumbers := numbers[:3]
	// fmt.Println(someNumbers)

	// speacialNumber := someNumbers[:1]
	// fmt.Println(speacialNumber)

	// speacialNumber[0] = 23.0
	// fmt.Println(numbers)
	// fmt.Println(someNumbers)

	number := []int{10, 20, 30, 40, 50} // Go infers the length as 5
	fmt.Println(number)
	fmt.Println(len(number))
	number = append(number, 60, 70, 80, 100)
	fmt.Println(number)
}
