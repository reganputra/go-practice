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

// create constructor function to initialize User struct

func NewUser(firstName, lastName, birthdate string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		createdAt: time.Now(),
	}
}

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var users *User

	//users = User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	Birthdate: birthdate,
	//	createdAt: time.Now(),
	//} // struct literal notation

	users = NewUser(firstName, lastName, birthdate) // constructor function

	users.outputDetail()
	users.remove()
	users.outputDetail()
}

func (user *User) outputDetail() {
	fmt.Println("Outputting user details...")
	fmt.Println("First Name:", user.FirstName)
	fmt.Println("Last Name:", user.LastName)
	fmt.Println("Birthdate:", user.Birthdate)
}

func (user *User) remove() {
	fmt.Println("Removing user...")
	user.FirstName = ""
	user.LastName = ""
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
