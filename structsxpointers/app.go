// package main

// import (
// 	"fmt"
// 	"time"
// )

// type User struct {
// 	FirstName string
// 	LastName  string
// 	Birthday  string
// 	Hobby     string
// 	AddedAt   time.Time
// }

// func main() {
// 	firstUser := User{
// 		FirstName: "John",
// 		LastName:  "Doe",
// 		Birthday:  "11/11/11",
// 		Hobby:     "football",
// 		AddedAt:   time.Now(),
// 	}

// 	// secondUser := User{"Lek", "Dan", "12/12/12", "cricket", time.Now()}

// 	// fmt.Println("Second user's first name:", secondUser.FirstName)
// 	// fmt.Println("First user's first name:", firstUser.FirstName)

// 	// updatedUser := getInputs(firstUser)
// 	// fmt.Printf("Updated user: %+v\n", updatedUser)
// 	// fmt.Println("Original first user's first name:", firstUser.FirstName)
// 	firstUser.FirstName = "jake"
// 	fmt.Println(firstUser.FirstName)
// }

// // func getInputs(user User) User {
// // 	fmt.Print("What is your first name? ")
// // 	fmt.Scan(&user.FirstName)

// // 	fmt.Print("What is your last name? ")
// // 	fmt.Scan(&user.LastName)

// // 	fmt.Print("What is your birthday? ")
// // 	fmt.Scan(&user.Birthday)

// // 	fmt.Print("What is your hobby? ")
// // 	fmt.Scan(&user.Hobby)

// // 	return user
// // }

package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthDate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	outputUserDetails(&appUser)
}

func outputUserDetails(u *user) {
	// ...

	fmt.Println(u.firstName, u.lastName, u.birthDate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
