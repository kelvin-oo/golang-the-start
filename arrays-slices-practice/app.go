package main

import "fmt"

type product struct {
	title string
	id    int
	price float64
}

func main() {
	// test one
	hobbies := [3]string{"football", "anime", "coding"}
	fmt.Println(hobbies)

	// second
	fmt.Println(hobbies[0])
	newHobbies := append([]string{}, hobbies[1], hobbies[2])
	fmt.Println(newHobbies)

	// three
	sliceHobbies := hobbies[:2]
	fmt.Println(sliceHobbies)

	// four
	fourthSlice := sliceHobbies[1:3]
	fmt.Println(fourthSlice)

	// fifth
	courseGoals := []string{"golang apis", "solid goalng"}
	fmt.Println(courseGoals)

	// sixth
	courseGoals[1] = "golang solid"
	fmt.Println(courseGoals)
	courseGoals = append(courseGoals, "golang syntax")
	fmt.Println(courseGoals)

	// seventh
	productList := []product{}
	fmt.Println("Initial productList:", productList) // Output: []
	product1 := product{title: "Laptop", id: 1, price: 999.99}
	product2 := product{title: "Smartphone", id: 2, price: 499.99}
	productList = append(productList, product1, product2)

	fmt.Println("Updated productList:", productList)
	product3 := product{title: "phone", id: 3, price: 49.9}
	productList = append(productList, product3)
	fmt.Println(productList)
}
