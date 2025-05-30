package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	createdAt time.Time
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var users User

	users = User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		createdAt: time.Now(),
	}

	fmt.Println(firstName, lastName, birthdate)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
