package main

import "fmt"

func main() {
	five := factorize(8)
	fmt.Println(five)
}

func factorize(number int) int {
	if number == 0 {
		return 1
	}
	return number * factorize(number-1)
	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result = result * i
	// }
	// return result
}
