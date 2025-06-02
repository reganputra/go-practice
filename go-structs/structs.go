package main

import (
	"fmt"
	"go-structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var users *user.User

	//users = User{
	//	FirstName: firstName,
	//	LastName:  lastName,
	//	Birthdate: birthdate,
	//	createdAt: time.Now(),
	//} // struct literal notation

	users, err := user.NewUser(firstName, lastName, birthdate) // constructor function

	if err != nil {
		fmt.Println("Error creating user:", err)
		return
	}

	users.OutputDetail()
	users.Remove()
	users.OutputDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
