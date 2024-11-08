package main

import "fmt"

func main() {

	var numbers = []int{1, 2, 3, 4, 5}
	result := sumAll(1, 2, 3, 4, 5)
	fmt.Println(result)
	anotherSum := sumAll(numbers...)
	fmt.Print(anotherSum)
}

func sumAll(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}
	return sum
}
