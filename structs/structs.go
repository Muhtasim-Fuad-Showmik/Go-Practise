package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user

	appUser = user{
		firstName: firstName,
		lastName:  lastName,
		birthdate: birthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	outputUserData(&appUser)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}

func outputUserData(outputUser *user) {
	fmt.Println("FirstName: ", outputUser.firstName)
	fmt.Println("LastName: ", outputUser.lastName)
	fmt.Println("Birthdate: ", outputUser.birthdate)
	fmt.Println("CreatedAt: ", outputUser.createdAt)
}
